// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: listing.proto

package v1

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
)

// Validate checks the field values on ID with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *ID) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return IDValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	return nil
}

// IDValidationError is the validation error returned by ID.Validate if the
// designated constraints aren't met.
type IDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IDValidationError) ErrorName() string { return "IDValidationError" }

// Error satisfies the builtin error interface
func (e IDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IDValidationError{}

// Validate checks the field values on UpdateReq with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UpdateReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return UpdateReqValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	if v, ok := interface{}(m.GetGame()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateReqValidationError{
				field:  "Game",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateReqValidationError is the validation error returned by
// UpdateReq.Validate if the designated constraints aren't met.
type UpdateReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateReqValidationError) ErrorName() string { return "UpdateReqValidationError" }

// Error satisfies the builtin error interface
func (e UpdateReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateReqValidationError{}

// Validate checks the field values on Game with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Game) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetId()) < 3 {
		return GameValidationError{
			field:  "Id",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetCategory()) < 3 {
		return GameValidationError{
			field:  "Category",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetTitle()) < 3 {
		return GameValidationError{
			field:  "Title",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetSubtitle()) < 3 {
		return GameValidationError{
			field:  "Subtitle",
			reason: "value length must be at least 3 runes",
		}
	}

	if utf8.RuneCountInString(m.GetDescription()) < 3 {
		return GameValidationError{
			field:  "Description",
			reason: "value length must be at least 3 runes",
		}
	}

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GameValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Type

	if len(m.GetTags()) > 0 {

		for idx, item := range m.GetTags() {
			_, _ = idx, item

			if utf8.RuneCountInString(item) < 3 {
				return GameValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "value length must be at least 3 runes",
				}
			}

		}

	}

	if utf8.RuneCountInString(m.GetAuthor()) < 3 {
		return GameValidationError{
			field:  "Author",
			reason: "value length must be at least 3 runes",
		}
	}

	if m.GetReplayBundleUrlJson() != "" {

		if uri, err := url.Parse(m.GetReplayBundleUrlJson()); err != nil {
			return GameValidationError{
				field:  "ReplayBundleUrlJson",
				reason: "value must be a valid URI",
				cause:  err,
			}
		} else if !uri.IsAbs() {
			return GameValidationError{
				field:  "ReplayBundleUrlJson",
				reason: "value must be absolute",
			}
		}

	}

	// no validation rules for Duration

	// no validation rules for IsDownloadable

	// no validation rules for IsStreamable

	if l := utf8.RuneCountInString(m.GetVersion()); l < 3 || l > 5 {
		return GameValidationError{
			field:  "Version",
			reason: "value length must be between 3 and 5 runes, inclusive",
		}
	}

	return nil
}

// GameValidationError is the validation error returned by Game.Validate if the
// designated constraints aren't met.
type GameValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GameValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GameValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GameValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GameValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GameValidationError) ErrorName() string { return "GameValidationError" }

// Error satisfies the builtin error interface
func (e GameValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGame.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GameValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GameValidationError{}

// Validate checks the field values on Images with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Images) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetId()); l < 3 || l > 20 {
		return ImagesValidationError{
			field:  "Id",
			reason: "value length must be between 3 and 20 runes, inclusive",
		}
	}

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		return ImagesValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return ImagesValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
	}

	// no validation rules for Type

	return nil
}

// ImagesValidationError is the validation error returned by Images.Validate if
// the designated constraints aren't met.
type ImagesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImagesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImagesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImagesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImagesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImagesValidationError) ErrorName() string { return "ImagesValidationError" }

// Error satisfies the builtin error interface
func (e ImagesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImages.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImagesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImagesValidationError{}

// Validate checks the field values on StreamGames with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StreamGames) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetListings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StreamGamesValidationError{
				field:  "Listings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// StreamGamesValidationError is the validation error returned by
// StreamGames.Validate if the designated constraints aren't met.
type StreamGamesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamGamesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamGamesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamGamesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamGamesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamGamesValidationError) ErrorName() string { return "StreamGamesValidationError" }

// Error satisfies the builtin error interface
func (e StreamGamesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamGames.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamGamesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamGamesValidationError{}

// Validate checks the field values on Games with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Games) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetListings() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GamesValidationError{
					field:  fmt.Sprintf("Listings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GamesValidationError is the validation error returned by Games.Validate if
// the designated constraints aren't met.
type GamesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GamesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GamesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GamesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GamesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GamesValidationError) ErrorName() string { return "GamesValidationError" }

// Error satisfies the builtin error interface
func (e GamesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGames.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GamesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GamesValidationError{}