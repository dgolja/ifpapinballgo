// Package types provides flexible JSON unmarshalling types for working with the
// IFPA (International Flipper Pinball Association) API, which returns fields in
// inconsistent formats — for example, integers encoded as strings, booleans as
// "Y"/"N", or absent values as empty strings instead of null.
//
// Long term plan is that IFPA will address the inconsistency in the OpenAPI spec
// so that all this workaround will not be necessary anymore.
package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Ptr returns a pointer to v. It is a convenience helper for passing literal
// values to functions or struct fields that expect a pointer.
//
// Example:
//
//	s := Ptr("hello") // s is *string pointing to "hello"
func Ptr[T any](v T) *T {
	return &v
}

// StringInt is an integer type that unmarshals from JSON integers, numeric
// strings, empty strings, or null. Empty strings and null are treated as 0.
//
// Accepted JSON inputs:
//
//	42      → 42
//	"42"    → 42
//	""      → 0
//	null    → 0
//	"abc"   → error
type StringInt int

// UnmarshalJSON implements [json.Unmarshaler]. It accepts a JSON number,
// a quoted numeric string, an empty string (→ 0), or null (→ 0).
func (f *StringInt) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*f = StringInt(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	parsed, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("cannot unmarshal %s into StringInt", string(data))
	}
	*f = StringInt(parsed)
	return nil
}

// String returns the decimal string representation of f.
func (f StringInt) String() string {
	return fmt.Sprintf("%d", f)
}

// StringFloat64 is a float64 type that unmarshals from JSON numbers, numeric
// strings, or null. Null is treated as 0. Empty strings are not accepted and
// will return an error.
//
// Accepted JSON inputs:
//
//	3.14    → 3.14
//	"3.14"  → 3.14
//	"10"    → 10.0
//	null    → 0
//	"abc"   → error
type StringFloat64 float64

// UnmarshalJSON implements [json.Unmarshaler]. It accepts a JSON number,
// a quoted numeric string, or null (→ 0).
func (f *StringFloat64) UnmarshalJSON(data []byte) error {
	var i float64
	if err := json.Unmarshal(data, &i); err == nil {
		*f = StringFloat64(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("cannot unmarshal %s into StringFloat64", string(data))
	}
	parsed, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = StringFloat64(parsed)
	return nil
}

// String returns the decimal string representation of f (e.g. "3.140000").
func (f StringFloat64) String() string {
	return fmt.Sprintf("%f", f)
}

// StringBool is a bool type that unmarshals from JSON booleans, the strings
// "Y"/"N", strconv-parseable bool strings ("true"/"false"/"1"/"0" etc.),
// empty strings, or null. Empty strings and null are treated as false.
//
// Accepted JSON inputs:
//
//	true      → true
//	false     → false
//	"Y"       → true
//	"N"       → false
//	"true"    → true
//	"false"   → false
//	""        → false
//	null      → false
//	"maybe"   → error
type StringBool bool

// UnmarshalJSON implements [json.Unmarshaler]. It accepts a JSON boolean,
// the strings "Y" (true) or "N" (false), any string accepted by
// [strconv.ParseBool], an empty string (→ false), or null (→ false).
func (f *StringBool) UnmarshalJSON(data []byte) error {
	var i bool
	if err := json.Unmarshal(data, &i); err == nil {
		*f = StringBool(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("cannot unmarshal %s into StringBool", string(data))
	}
	if s == "" {
		return nil
	}
	if s == "N" {
		*f = StringBool(false)
		return nil
	}
	if s == "Y" {
		*f = StringBool(true)
		return nil
	}
	parsed, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*f = StringBool(parsed)
	return nil
}

// String returns "true" or "false".
func (f StringBool) String() string {
	return fmt.Sprintf("%t", f)
}

// FlexibleInt represents a JSON field that may be an integer, a numeric string,
// an empty string, or null. It distinguishes between "no value" and the integer
// zero, which plain *int does not when dealing with empty-string sentinels.
//
// Accepted JSON inputs and resulting state:
//
//	57      → {Value: 57, IsEmpty: false}
//	"57"    → {Value: 57, IsEmpty: false}
//	0       → {Value: 0,  IsEmpty: false}
//	""      → {Value: 0,  IsEmpty: true}
//	null    → {Value: 0,  IsEmpty: true}
//	"abc"   → error
//
// Use [FlexibleInt.HasValue] to test whether a value is present, and
// [FlexibleInt.GetValue] to retrieve it.
type FlexibleInt struct {
	Value   int
	IsEmpty bool
}

// UnmarshalJSON implements [json.Unmarshaler]. It accepts a JSON integer,
// a quoted numeric string, an empty string (→ IsEmpty), or null (→ IsEmpty).
// Non-numeric strings return an error.
func (fi *FlexibleInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		fi.IsEmpty = true
		fi.Value = 0
		return nil
	}

	// Try to unmarshal as integer first
	var num int
	if err := json.Unmarshal(data, &num); err == nil {
		fi.Value = num
		fi.IsEmpty = false
		return nil
	}

	// Try to unmarshal as string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str == "" {
			fi.IsEmpty = true
			fi.Value = 0
		} else {
			// Try to parse string as integer
			if val, err := strconv.Atoi(str); err == nil {
				fi.Value = val
				fi.IsEmpty = false
			} else {
				fi.IsEmpty = true
				fi.Value = 0
				return fmt.Errorf("UnmarshalJSON: unable to parse '%s'", str)
			}
		}
		return nil
	}

	return fmt.Errorf("cannot unmarshal %s into FlexibleInt", string(data))
}

// MarshalJSON implements [json.Marshaler]. An empty FlexibleInt marshals to
// the JSON string "" to round-trip through the IFPA API format; a non-empty
// value marshals to a plain JSON integer.
func (fi FlexibleInt) MarshalJSON() ([]byte, error) {
	if fi.IsEmpty {
		return json.Marshal("")
	}
	return json.Marshal(fi.Value)
}

// GetValue returns the integer value. Returns 0 when IsEmpty is true.
// Prefer checking [FlexibleInt.HasValue] before calling this method.
func (fi FlexibleInt) GetValue() int {
	return fi.Value
}

// HasValue reports whether the field contains a valid integer value
// (i.e. was not null or an empty string in the source JSON).
func (fi FlexibleInt) HasValue() bool {
	return !fi.IsEmpty
}

// TourRelatedWinner represents the winner field on tournament-related responses.
// The IFPA API returns either an empty string (no winner yet) or a player
// object. All fields are nil when the winner is not yet determined.
//
// Note: ideally the IFPA API would return {} instead of "" when unset.
type TourRelatedWinner struct {
	PlayerId    *StringInt `json:"player_id,omitempty"`
	Name        *string    `json:"name,omitempty"`
	CountryName *string    `json:"country_name,omitempty"`
	CountryCode *string    `json:"country_code,omitempty"`
}

// UnmarshalJSON implements [json.Unmarshaler]. An empty JSON string leaves all
// fields nil (no winner). A JSON object is decoded into the struct fields.
func (w *TourRelatedWinner) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		return nil // empty string — leave all fields nil
	}
	type Alias TourRelatedWinner
	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*w = TourRelatedWinner(alias)
	return nil
}
