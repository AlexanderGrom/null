package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
)

var (
	_ sql.Scanner              = (*String)(nil)
	_ driver.Valuer            = (*String)(nil)
	_ json.Marshaler           = (*String)(nil)
	_ json.Unmarshaler         = (*String)(nil)
	_ encoding.TextMarshaler   = (*String)(nil)
	_ encoding.TextUnmarshaler = (*String)(nil)
)

// NullBytes is a byte slice of JSON null
var NullBytes = []byte("null")

// String is a nullable string. String provides SQL and JSON serialization.
type String struct {
	sql.NullString
}

// NewString creates a new String
func NewString(s string, v bool) String {
	return String{
		sql.NullString{s, v},
	}
}

// StringFrom creates a new valid String from raw string.
func StringFrom(s string) String {
	return NewString(s, true)
}

// MarshalJSON implements json.Marshaler.
func (e String) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.String)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *String) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.String, e.Valid = "", false
		return nil
	}
	err = json.Unmarshal(data, &e.String)
	e.String = e.String
	e.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e String) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(e.String), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *String) UnmarshalText(text []byte) error {
	e.String = string(text)
	e.Valid = text != nil
	return nil
}

// IsZero returns true for zero value.
func (e String) IsZero() bool {
	return len(e.String) == 0
}

// IsSet returns true for valid value.
func (e String) IsSet() bool {
	return e.Valid
}
