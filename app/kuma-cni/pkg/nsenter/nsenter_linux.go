package nsenter

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

// Filesystems constants.
const (
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/magic.h
	nsFSMagic   = 0x6e736673
	procFSMagic = 0x9fa0

	procRootPath = "/proc"
	nsDirPath    = "ns"
	taskDirPath  = "task"
)

type nsPair struct {
	targetNS *os.File
	threadNS *os.File
}

func init() {
	var ns = map[NSType]int{
		NSTypeCGroup: unix.CLONE_NEWCGROUP,
		NSTypeIPC:    unix.CLONE_NEWIPC,
		NSTypeNet:    unix.CLONE_NEWNET,
		NSTypePID:    unix.CLONE_NEWPID,
		NSTypeUTS:    unix.CLONE_NEWUTS,
	}

	for k, v := range ns {
		if _, err := os.Stat(fmt.Sprint("/proc/self/ns/", string(k))); err == nil {
			CloneFlagsTable[k] = v
		}
	}
}

func getNSPathFromPID(pid int, nsType NSType) string {
	return filepath.Join(procRootPath, strconv.Itoa(pid), nsDirPath, string(nsType))
}

func getCurrentThreadNSPath(nsType NSType) string {
	return filepath.Join(procRootPath, strconv.Itoa(os.Getpid()),
		taskDirPath, strconv.Itoa(unix.Gettid()), nsDirPath, string(nsType))
}

func setNS(nsFile *os.File, nsType NSType) error {
	if nsFile == nil {
		return fmt.Errorf("File handler cannot be nil")
	}

	nsFlag, exist := CloneFlagsTable[nsType]
	if !exist {
		return fmt.Errorf("Unknown namespace type %q", nsType)
	}

	if err := unix.Setns(int(nsFile.Fd()), nsFlag); err != nil {
		return fmt.Errorf("Error switching to ns %v: %v", nsFile.Name(), err)
	}

	return nil
}

// getFileFromNS checks the provided file path actually matches a real
// namespace filesystem, and then opens it to return a handler to this
// file. This is needed since the system call setns() expects a file
// descriptor to enter the given namespace.
func getFileFromNS(nsPath string) (*os.File, error) {
	stat := syscall.Statfs_t{}
	if err := syscall.Statfs(nsPath, &stat); err != nil {
		return nil, fmt.Errorf("failed to Statfs %q: %v", nsPath, err)
	}

	switch stat.Type {
	case nsFSMagic, procFSMagic:
		break
	default:
		return nil, fmt.Errorf("unknown FS magic on %q: %x", nsPath, stat.Type)
	}

	file, err := os.Open(nsPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// NsEnter executes the passed closure under the given namespace,
// restoring the original namespace afterwards.
func NsEnter(nsList []Namespace, toRun func() error) error {
	targetNSList := make(map[NSType]*nsPair)

	// Open all targeted namespaces.
	for _, ns := range nsList {
		targetNSPath := ns.Path
		if targetNSPath == "" {
			targetNSPath = getNSPathFromPID(ns.PID, ns.Type)
		}

		targetNS, err := getFileFromNS(targetNSPath)
		if err != nil {
			return fmt.Errorf("failed to open target ns: %v", err)
		}
		defer targetNS.Close()

		targetNSList[ns.Type] = &nsPair{
			targetNS: targetNS,
		}
	}

	containedCall := func() error {
		for nsType := range targetNSList {
			threadNS, err := getFileFromNS(getCurrentThreadNSPath(nsType))
			if err != nil {
				return fmt.Errorf("failed to open current ns: %v", err)
			}
			defer threadNS.Close()

			targetNSList[nsType].threadNS = threadNS
		}

		// Switch to namespaces all at once.
		for nsType, pair := range targetNSList {
			// Switch to targeted namespace.
			if err := setNS(pair.targetNS, nsType); err != nil {
				return fmt.Errorf("error switching to ns %v: %v", pair.targetNS.Name(), err)
			}
			// Switch back to initial namespace after closure return.
			threadNS := pair.threadNS
			_nsType := nsType
			defer func() {
				_ = setNS(threadNS, _nsType)
			}()
		}

		return toRun()
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var innerError error
	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
		innerError = containedCall()
	}()
	wg.Wait()

	return innerError
}
