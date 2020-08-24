package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"strconv"
)

var (
	_ sql.Scanner              = (*Float64)(nil)
	_ driver.Valuer            = (*Float64)(nil)
	_ json.Marshaler           = (*Float64)(nil)
	_ json.Unmarshaler         = (*Float64)(nil)
	_ encoding.TextMarshaler   = (*Float64)(nil)
	_ encoding.TextUnmarshaler = (*Float64)(nil)
)

// Float64 is a nullable float64. Float64 provides SQL and JSON serialization.
type Float64 struct {
	sql.NullFloat64
}

// NewFloat64 creates a new Float64
func NewFloat64(f float64, v bool) Float64 {
	return Float64{
		sql.NullFloat64{f, v},
	}
}

// Float64From creates a new valid Float64 from raw float64.
func Float64From(f float64) Float64 {
	return NewFloat64(f, true)
}

// MarshalJSON implements json.Marshaler.
func (e Float64) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.Float64)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Float64) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.Float64, e.Valid = 0, false
		return nil
	}
	err = json.Unmarshal(data, &e.Float64)
	e.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e Float64) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatFloat(e.Float64, 'f', -1, 64)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Float64) UnmarshalText(text []byte) (err error) {
	e.Float64, err = strconv.ParseFloat(string(text), 64)
	e.Valid = err == nil
	return nil
}

// IsZero returns true for zero value.
func (e Float64) IsZero() bool {
	return e.Float64 == 0
}

// IsSet returns true for valid value.
func (e Float64) IsSet() bool {
	return e.Valid
}

// IsSetZero returns true if value is set to zero.
func (e Float64) IsSetZero() bool {
	return e.IsSet() && e.IsZero()
}
