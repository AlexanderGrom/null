package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"errors"
)

var (
	_ sql.Scanner              = (*Bool)(nil)
	_ driver.Valuer            = (*Bool)(nil)
	_ json.Marshaler           = (*Bool)(nil)
	_ json.Unmarshaler         = (*Bool)(nil)
	_ encoding.TextMarshaler   = (*Bool)(nil)
	_ encoding.TextUnmarshaler = (*Bool)(nil)
)

// Bool is a nullable bool. Bool provides SQL and JSON serialization.
type Bool struct {
	sql.NullBool
}

// NewBool creates a new Bool
func NewBool(b bool, v bool) Bool {
	return Bool{
		sql.NullBool{b, v},
	}
}

// BoolFrom creates a new valid Bool from raw bool.
func BoolFrom(b bool) Bool {
	return NewBool(b, true)
}

// MarshalJSON implements json.Marshaler.
func (e Bool) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.Bool)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Bool) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.Bool, e.Valid = false, false
		return nil
	}
	err = json.Unmarshal(data, &e.Bool)
	e.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e Bool) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	if !e.Bool {
		return []byte("false"), nil
	}
	return []byte("true"), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Bool) UnmarshalText(text []byte) error {
	var val = string(text)
	switch val {
	case "true":
		e.Bool, e.Valid = true, true
	case "false":
		e.Bool, e.Valid = false, true
	default:
		e.Bool, e.Valid = false, false
		return errors.New(`invalid bool: "` + val + `"`)
	}
	return nil
}

// IsZero returns true for zero value.
func (e *Bool) IsZero() bool {
	return !e.Bool
}

// IsSet returns true for valid value.
func (e *Bool) IsSet() bool {
	return e.Valid
}
