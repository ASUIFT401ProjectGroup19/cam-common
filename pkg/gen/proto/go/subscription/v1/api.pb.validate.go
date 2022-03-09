// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: subscription/v1/api.proto

package subscriptionv1

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

// Validate checks the field values on FollowRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FollowRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FollowRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FollowRequestMultiError, or
// nil if none found.
func (m *FollowRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FollowRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return FollowRequestMultiError(errors)
	}
	return nil
}

// FollowRequestMultiError is an error wrapping multiple validation errors
// returned by FollowRequest.ValidateAll() if the designated constraints
// aren't met.
type FollowRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FollowRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FollowRequestMultiError) AllErrors() []error { return m }

// FollowRequestValidationError is the validation error returned by
// FollowRequest.Validate if the designated constraints aren't met.
type FollowRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FollowRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FollowRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FollowRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FollowRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FollowRequestValidationError) ErrorName() string { return "FollowRequestValidationError" }

// Error satisfies the builtin error interface
func (e FollowRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFollowRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FollowRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FollowRequestValidationError{}

// Validate checks the field values on FollowResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FollowResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FollowResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FollowResponseMultiError,
// or nil if none found.
func (m *FollowResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FollowResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FollowResponseMultiError(errors)
	}
	return nil
}

// FollowResponseMultiError is an error wrapping multiple validation errors
// returned by FollowResponse.ValidateAll() if the designated constraints
// aren't met.
type FollowResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FollowResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FollowResponseMultiError) AllErrors() []error { return m }

// FollowResponseValidationError is the validation error returned by
// FollowResponse.Validate if the designated constraints aren't met.
type FollowResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FollowResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FollowResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FollowResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FollowResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FollowResponseValidationError) ErrorName() string { return "FollowResponseValidationError" }

// Error satisfies the builtin error interface
func (e FollowResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFollowResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FollowResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FollowResponseValidationError{}

// Validate checks the field values on UnfollowRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnfollowRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnfollowRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnfollowRequestMultiError, or nil if none found.
func (m *UnfollowRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UnfollowRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return UnfollowRequestMultiError(errors)
	}
	return nil
}

// UnfollowRequestMultiError is an error wrapping multiple validation errors
// returned by UnfollowRequest.ValidateAll() if the designated constraints
// aren't met.
type UnfollowRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnfollowRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnfollowRequestMultiError) AllErrors() []error { return m }

// UnfollowRequestValidationError is the validation error returned by
// UnfollowRequest.Validate if the designated constraints aren't met.
type UnfollowRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnfollowRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnfollowRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnfollowRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnfollowRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnfollowRequestValidationError) ErrorName() string { return "UnfollowRequestValidationError" }

// Error satisfies the builtin error interface
func (e UnfollowRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnfollowRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnfollowRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnfollowRequestValidationError{}

// Validate checks the field values on UnfollowResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UnfollowResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnfollowResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnfollowResponseMultiError, or nil if none found.
func (m *UnfollowResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UnfollowResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UnfollowResponseMultiError(errors)
	}
	return nil
}

// UnfollowResponseMultiError is an error wrapping multiple validation errors
// returned by UnfollowResponse.ValidateAll() if the designated constraints
// aren't met.
type UnfollowResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnfollowResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnfollowResponseMultiError) AllErrors() []error { return m }

// UnfollowResponseValidationError is the validation error returned by
// UnfollowResponse.Validate if the designated constraints aren't met.
type UnfollowResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnfollowResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnfollowResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnfollowResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnfollowResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnfollowResponseValidationError) ErrorName() string { return "UnfollowResponseValidationError" }

// Error satisfies the builtin error interface
func (e UnfollowResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnfollowResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnfollowResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnfollowResponseValidationError{}
