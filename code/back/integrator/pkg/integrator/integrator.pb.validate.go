// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: integrator.proto

package integrator

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Specification with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Specification) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Specification with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SpecificationMultiError, or
// nil if none found.
func (m *Specification) ValidateAll() error {
	return m.validate(true)
}

func (m *Specification) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Content

	// no validation rules for GitPath

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SpecificationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SpecificationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpecificationValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SpecificationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SpecificationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SpecificationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SpecificationMultiError(errors)
	}

	return nil
}

// SpecificationMultiError is an error wrapping multiple validation errors
// returned by Specification.ValidateAll() if the designated constraints
// aren't met.
type SpecificationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SpecificationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SpecificationMultiError) AllErrors() []error { return m }

// SpecificationValidationError is the validation error returned by
// Specification.Validate if the designated constraints aren't met.
type SpecificationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SpecificationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SpecificationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SpecificationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SpecificationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SpecificationValidationError) ErrorName() string { return "SpecificationValidationError" }

// Error satisfies the builtin error interface
func (e SpecificationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSpecification.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SpecificationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SpecificationValidationError{}

// Validate checks the field values on CommitPushResult with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CommitPushResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommitPushResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommitPushResultMultiError, or nil if none found.
func (m *CommitPushResult) ValidateAll() error {
	return m.validate(true)
}

func (m *CommitPushResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CommitHash

	// no validation rules for IsSuccess

	if len(errors) > 0 {
		return CommitPushResultMultiError(errors)
	}

	return nil
}

// CommitPushResultMultiError is an error wrapping multiple validation errors
// returned by CommitPushResult.ValidateAll() if the designated constraints
// aren't met.
type CommitPushResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommitPushResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommitPushResultMultiError) AllErrors() []error { return m }

// CommitPushResultValidationError is the validation error returned by
// CommitPushResult.Validate if the designated constraints aren't met.
type CommitPushResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommitPushResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommitPushResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommitPushResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommitPushResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommitPushResultValidationError) ErrorName() string { return "CommitPushResultValidationError" }

// Error satisfies the builtin error interface
func (e CommitPushResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommitPushResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommitPushResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommitPushResultValidationError{}

// Validate checks the field values on MLDevResult with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MLDevResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MLDevResult with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MLDevResultMultiError, or
// nil if none found.
func (m *MLDevResult) ValidateAll() error {
	return m.validate(true)
}

func (m *MLDevResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsSuccess

	if len(errors) > 0 {
		return MLDevResultMultiError(errors)
	}

	return nil
}

// MLDevResultMultiError is an error wrapping multiple validation errors
// returned by MLDevResult.ValidateAll() if the designated constraints aren't met.
type MLDevResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MLDevResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MLDevResultMultiError) AllErrors() []error { return m }

// MLDevResultValidationError is the validation error returned by
// MLDevResult.Validate if the designated constraints aren't met.
type MLDevResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MLDevResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MLDevResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MLDevResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MLDevResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MLDevResultValidationError) ErrorName() string { return "MLDevResultValidationError" }

// Error satisfies the builtin error interface
func (e MLDevResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMLDevResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MLDevResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MLDevResultValidationError{}

// Validate checks the field values on GetSpecificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSpecificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSpecificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSpecificationRequestMultiError, or nil if none found.
func (m *GetSpecificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSpecificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetSpecificationRequestMultiError(errors)
	}

	return nil
}

// GetSpecificationRequestMultiError is an error wrapping multiple validation
// errors returned by GetSpecificationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetSpecificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSpecificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSpecificationRequestMultiError) AllErrors() []error { return m }

// GetSpecificationRequestValidationError is the validation error returned by
// GetSpecificationRequest.Validate if the designated constraints aren't met.
type GetSpecificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSpecificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSpecificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSpecificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSpecificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSpecificationRequestValidationError) ErrorName() string {
	return "GetSpecificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSpecificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSpecificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSpecificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSpecificationRequestValidationError{}

// Validate checks the field values on SaveSpecificationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SaveSpecificationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SaveSpecificationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SaveSpecificationRequestMultiError, or nil if none found.
func (m *SaveSpecificationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SaveSpecificationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetSpecification()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SaveSpecificationRequestValidationError{
					field:  "Specification",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SaveSpecificationRequestValidationError{
					field:  "Specification",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSpecification()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SaveSpecificationRequestValidationError{
				field:  "Specification",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SaveSpecificationRequestMultiError(errors)
	}

	return nil
}

// SaveSpecificationRequestMultiError is an error wrapping multiple validation
// errors returned by SaveSpecificationRequest.ValidateAll() if the designated
// constraints aren't met.
type SaveSpecificationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SaveSpecificationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SaveSpecificationRequestMultiError) AllErrors() []error { return m }

// SaveSpecificationRequestValidationError is the validation error returned by
// SaveSpecificationRequest.Validate if the designated constraints aren't met.
type SaveSpecificationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SaveSpecificationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SaveSpecificationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SaveSpecificationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SaveSpecificationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SaveSpecificationRequestValidationError) ErrorName() string {
	return "SaveSpecificationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SaveSpecificationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSaveSpecificationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SaveSpecificationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SaveSpecificationRequestValidationError{}

// Validate checks the field values on RunMLDevRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RunMLDevRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RunMLDevRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RunMLDevRequestMultiError, or nil if none found.
func (m *RunMLDevRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RunMLDevRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetPath()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RunMLDevRequestValidationError{
					field:  "Path",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RunMLDevRequestValidationError{
					field:  "Path",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPath()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RunMLDevRequestValidationError{
				field:  "Path",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RunMLDevRequestMultiError(errors)
	}

	return nil
}

// RunMLDevRequestMultiError is an error wrapping multiple validation errors
// returned by RunMLDevRequest.ValidateAll() if the designated constraints
// aren't met.
type RunMLDevRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RunMLDevRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RunMLDevRequestMultiError) AllErrors() []error { return m }

// RunMLDevRequestValidationError is the validation error returned by
// RunMLDevRequest.Validate if the designated constraints aren't met.
type RunMLDevRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RunMLDevRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RunMLDevRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RunMLDevRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RunMLDevRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RunMLDevRequestValidationError) ErrorName() string { return "RunMLDevRequestValidationError" }

// Error satisfies the builtin error interface
func (e RunMLDevRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRunMLDevRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RunMLDevRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RunMLDevRequestValidationError{}

// Validate checks the field values on MLDevPath with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MLDevPath) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MLDevPath with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MLDevPathMultiError, or nil
// if none found.
func (m *MLDevPath) ValidateAll() error {
	return m.validate(true)
}

func (m *MLDevPath) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	if len(errors) > 0 {
		return MLDevPathMultiError(errors)
	}

	return nil
}

// MLDevPathMultiError is an error wrapping multiple validation errors returned
// by MLDevPath.ValidateAll() if the designated constraints aren't met.
type MLDevPathMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MLDevPathMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MLDevPathMultiError) AllErrors() []error { return m }

// MLDevPathValidationError is the validation error returned by
// MLDevPath.Validate if the designated constraints aren't met.
type MLDevPathValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MLDevPathValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MLDevPathValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MLDevPathValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MLDevPathValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MLDevPathValidationError) ErrorName() string { return "MLDevPathValidationError" }

// Error satisfies the builtin error interface
func (e MLDevPathValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMLDevPath.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MLDevPathValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MLDevPathValidationError{}

// Validate checks the field values on GetHelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetHelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHelloRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHelloRequestMultiError, or nil if none found.
func (m *GetHelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetHelloRequestMultiError(errors)
	}

	return nil
}

// GetHelloRequestMultiError is an error wrapping multiple validation errors
// returned by GetHelloRequest.ValidateAll() if the designated constraints
// aren't met.
type GetHelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHelloRequestMultiError) AllErrors() []error { return m }

// GetHelloRequestValidationError is the validation error returned by
// GetHelloRequest.Validate if the designated constraints aren't met.
type GetHelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHelloRequestValidationError) ErrorName() string { return "GetHelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetHelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHelloRequestValidationError{}

// Validate checks the field values on GetHelloResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetHelloResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetHelloResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetHelloResponseMultiError, or nil if none found.
func (m *GetHelloResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetHelloResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pong

	if len(errors) > 0 {
		return GetHelloResponseMultiError(errors)
	}

	return nil
}

// GetHelloResponseMultiError is an error wrapping multiple validation errors
// returned by GetHelloResponse.ValidateAll() if the designated constraints
// aren't met.
type GetHelloResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetHelloResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetHelloResponseMultiError) AllErrors() []error { return m }

// GetHelloResponseValidationError is the validation error returned by
// GetHelloResponse.Validate if the designated constraints aren't met.
type GetHelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHelloResponseValidationError) ErrorName() string { return "GetHelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetHelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHelloResponseValidationError{}
