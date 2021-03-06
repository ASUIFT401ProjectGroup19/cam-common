// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: gallery/v1/api.proto

package galleryv1

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

// Validate checks the field values on GalleryRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GalleryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GalleryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GalleryRequestMultiError,
// or nil if none found.
func (m *GalleryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GalleryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for BatchSize

	// no validation rules for UserId

	if len(errors) > 0 {
		return GalleryRequestMultiError(errors)
	}
	return nil
}

// GalleryRequestMultiError is an error wrapping multiple validation errors
// returned by GalleryRequest.ValidateAll() if the designated constraints
// aren't met.
type GalleryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GalleryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GalleryRequestMultiError) AllErrors() []error { return m }

// GalleryRequestValidationError is the validation error returned by
// GalleryRequest.Validate if the designated constraints aren't met.
type GalleryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GalleryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GalleryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GalleryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GalleryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GalleryRequestValidationError) ErrorName() string { return "GalleryRequestValidationError" }

// Error satisfies the builtin error interface
func (e GalleryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGalleryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GalleryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GalleryRequestValidationError{}

// Validate checks the field values on GalleryResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GalleryResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GalleryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GalleryResponseMultiError, or nil if none found.
func (m *GalleryResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GalleryResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GalleryResponseValidationError{
						field:  fmt.Sprintf("Posts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GalleryResponseValidationError{
						field:  fmt.Sprintf("Posts[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GalleryResponseValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GalleryResponseMultiError(errors)
	}
	return nil
}

// GalleryResponseMultiError is an error wrapping multiple validation errors
// returned by GalleryResponse.ValidateAll() if the designated constraints
// aren't met.
type GalleryResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GalleryResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GalleryResponseMultiError) AllErrors() []error { return m }

// GalleryResponseValidationError is the validation error returned by
// GalleryResponse.Validate if the designated constraints aren't met.
type GalleryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GalleryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GalleryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GalleryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GalleryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GalleryResponseValidationError) ErrorName() string { return "GalleryResponseValidationError" }

// Error satisfies the builtin error interface
func (e GalleryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGalleryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GalleryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GalleryResponseValidationError{}
