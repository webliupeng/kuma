package rbac_test

import (
	"testing"

	"github.com/kumahq/kuma/pkg/test"
)

func TestManager(t *testing.T) {
	test.RunSpecs(t, "RBAC")
}
