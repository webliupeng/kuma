// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mesh/v1alpha1/mesh.proto

package v1alpha1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _mesh_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Mesh with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Mesh) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMtls()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MeshValidationError{
				field:  "Mtls",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTracing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MeshValidationError{
				field:  "Tracing",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLogging()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MeshValidationError{
				field:  "Logging",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetrics()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MeshValidationError{
				field:  "Metrics",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MeshValidationError is the validation error returned by Mesh.Validate if the
// designated constraints aren't met.
type MeshValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MeshValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MeshValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MeshValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MeshValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MeshValidationError) ErrorName() string { return "MeshValidationError" }

// Error satisfies the builtin error interface
func (e MeshValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMesh.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MeshValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MeshValidationError{}

// Validate checks the field values on CertificateAuthorityBackend with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CertificateAuthorityBackend) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Type

	if v, ok := interface{}(m.GetDpCert()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateAuthorityBackendValidationError{
				field:  "DpCert",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConf()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateAuthorityBackendValidationError{
				field:  "Conf",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CertificateAuthorityBackendValidationError is the validation error returned
// by CertificateAuthorityBackend.Validate if the designated constraints
// aren't met.
type CertificateAuthorityBackendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateAuthorityBackendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateAuthorityBackendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateAuthorityBackendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateAuthorityBackendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateAuthorityBackendValidationError) ErrorName() string {
	return "CertificateAuthorityBackendValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateAuthorityBackendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateAuthorityBackend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateAuthorityBackendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateAuthorityBackendValidationError{}

// Validate checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tracing) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DefaultBackend

	for idx, item := range m.GetBackends() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TracingValidationError{
					field:  fmt.Sprintf("Backends[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TracingValidationError is the validation error returned by Tracing.Validate
// if the designated constraints aren't met.
type TracingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TracingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TracingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TracingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TracingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TracingValidationError) ErrorName() string { return "TracingValidationError" }

// Error satisfies the builtin error interface
func (e TracingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TracingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TracingValidationError{}

// Validate checks the field values on TracingBackend with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TracingBackend) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetSampling()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TracingBackendValidationError{
				field:  "Sampling",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	if v, ok := interface{}(m.GetConf()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TracingBackendValidationError{
				field:  "Conf",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// TracingBackendValidationError is the validation error returned by
// TracingBackend.Validate if the designated constraints aren't met.
type TracingBackendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TracingBackendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TracingBackendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TracingBackendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TracingBackendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TracingBackendValidationError) ErrorName() string { return "TracingBackendValidationError" }

// Error satisfies the builtin error interface
func (e TracingBackendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracingBackend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TracingBackendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TracingBackendValidationError{}

// Validate checks the field values on ZipkinTracingBackendConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ZipkinTracingBackendConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Url

	// no validation rules for TraceId128Bit

	// no validation rules for ApiVersion

	return nil
}

// ZipkinTracingBackendConfigValidationError is the validation error returned
// by ZipkinTracingBackendConfig.Validate if the designated constraints aren't met.
type ZipkinTracingBackendConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZipkinTracingBackendConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZipkinTracingBackendConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZipkinTracingBackendConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZipkinTracingBackendConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZipkinTracingBackendConfigValidationError) ErrorName() string {
	return "ZipkinTracingBackendConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ZipkinTracingBackendConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZipkinTracingBackendConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZipkinTracingBackendConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZipkinTracingBackendConfigValidationError{}

// Validate checks the field values on Logging with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Logging) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DefaultBackend

	for idx, item := range m.GetBackends() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoggingValidationError{
					field:  fmt.Sprintf("Backends[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LoggingValidationError is the validation error returned by Logging.Validate
// if the designated constraints aren't met.
type LoggingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoggingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoggingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoggingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoggingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoggingValidationError) ErrorName() string { return "LoggingValidationError" }

// Error satisfies the builtin error interface
func (e LoggingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogging.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoggingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoggingValidationError{}

// Validate checks the field values on LoggingBackend with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoggingBackend) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Format

	// no validation rules for Type

	if v, ok := interface{}(m.GetConf()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoggingBackendValidationError{
				field:  "Conf",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LoggingBackendValidationError is the validation error returned by
// LoggingBackend.Validate if the designated constraints aren't met.
type LoggingBackendValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoggingBackendValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoggingBackendValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoggingBackendValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoggingBackendValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoggingBackendValidationError) ErrorName() string { return "LoggingBackendValidationError" }

// Error satisfies the builtin error interface
func (e LoggingBackendValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoggingBackend.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoggingBackendValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoggingBackendValidationError{}

// Validate checks the field values on FileLoggingBackendConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FileLoggingBackendConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Path

	return nil
}

// FileLoggingBackendConfigValidationError is the validation error returned by
// FileLoggingBackendConfig.Validate if the designated constraints aren't met.
type FileLoggingBackendConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileLoggingBackendConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileLoggingBackendConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileLoggingBackendConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileLoggingBackendConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileLoggingBackendConfigValidationError) ErrorName() string {
	return "FileLoggingBackendConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FileLoggingBackendConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileLoggingBackendConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileLoggingBackendConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileLoggingBackendConfigValidationError{}

// Validate checks the field values on TcpLoggingBackendConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TcpLoggingBackendConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Address

	return nil
}

// TcpLoggingBackendConfigValidationError is the validation error returned by
// TcpLoggingBackendConfig.Validate if the designated constraints aren't met.
type TcpLoggingBackendConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpLoggingBackendConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpLoggingBackendConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpLoggingBackendConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpLoggingBackendConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpLoggingBackendConfigValidationError) ErrorName() string {
	return "TcpLoggingBackendConfigValidationError"
}

// Error satisfies the builtin error interface
func (e TcpLoggingBackendConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpLoggingBackendConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpLoggingBackendConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpLoggingBackendConfigValidationError{}

// Validate checks the field values on Mesh_Mtls with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Mesh_Mtls) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EnabledBackend

	for idx, item := range m.GetBackends() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Mesh_MtlsValidationError{
					field:  fmt.Sprintf("Backends[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// Mesh_MtlsValidationError is the validation error returned by
// Mesh_Mtls.Validate if the designated constraints aren't met.
type Mesh_MtlsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Mesh_MtlsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Mesh_MtlsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Mesh_MtlsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Mesh_MtlsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Mesh_MtlsValidationError) ErrorName() string { return "Mesh_MtlsValidationError" }

// Error satisfies the builtin error interface
func (e Mesh_MtlsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMesh_Mtls.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Mesh_MtlsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Mesh_MtlsValidationError{}

// Validate checks the field values on CertificateAuthorityBackend_DpCert with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CertificateAuthorityBackend_DpCert) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRotation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CertificateAuthorityBackend_DpCertValidationError{
				field:  "Rotation",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CertificateAuthorityBackend_DpCertValidationError is the validation error
// returned by CertificateAuthorityBackend_DpCert.Validate if the designated
// constraints aren't met.
type CertificateAuthorityBackend_DpCertValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateAuthorityBackend_DpCertValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateAuthorityBackend_DpCertValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateAuthorityBackend_DpCertValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateAuthorityBackend_DpCertValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateAuthorityBackend_DpCertValidationError) ErrorName() string {
	return "CertificateAuthorityBackend_DpCertValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateAuthorityBackend_DpCertValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateAuthorityBackend_DpCert.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateAuthorityBackend_DpCertValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateAuthorityBackend_DpCertValidationError{}

// Validate checks the field values on
// CertificateAuthorityBackend_DpCert_Rotation with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CertificateAuthorityBackend_DpCert_Rotation) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Expiration

	return nil
}

// CertificateAuthorityBackend_DpCert_RotationValidationError is the validation
// error returned by CertificateAuthorityBackend_DpCert_Rotation.Validate if
// the designated constraints aren't met.
type CertificateAuthorityBackend_DpCert_RotationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) ErrorName() string {
	return "CertificateAuthorityBackend_DpCert_RotationValidationError"
}

// Error satisfies the builtin error interface
func (e CertificateAuthorityBackend_DpCert_RotationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCertificateAuthorityBackend_DpCert_Rotation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CertificateAuthorityBackend_DpCert_RotationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CertificateAuthorityBackend_DpCert_RotationValidationError{}
