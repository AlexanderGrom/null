package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"time"
)

var (
	_ sql.Scanner              = (*Time)(nil)
	_ driver.Valuer            = (*Time)(nil)
	_ json.Marshaler           = (*Time)(nil)
	_ json.Unmarshaler         = (*Time)(nil)
	_ encoding.TextMarshaler   = (*Time)(nil)
	_ encoding.TextUnmarshaler = (*Time)(nil)
)

// Time is a nullable string. Time provides SQL and JSON serialization.
type Time struct {
	sql.NullTime
}

// NewTime creates a new Time
func NewTime(t time.Time, v bool) Time {
	return Time{
		sql.NullTime{t, v},
	}
}

// TimeFrom creates a new valid Time from raw time.Time.
func TimeFrom(t time.Time) Time {
	return NewTime(t, true)
}

// MarshalJSON implements json.Marshaler.
func (e Time) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.Time)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Time) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.Time, e.Valid = time.Time{}, false
		return nil
	}
	err = json.Unmarshal(data, &e.Time)
	e.Time = e.Time
	e.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e Time) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return e.Time.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Time) UnmarshalText(text []byte) (err error) {
	err = e.Time.UnmarshalText(text)
	e.Valid = err == nil
	return err
}

// IsZero returns true for zero value.
func (e Time) IsZero() bool {
	return e.Time.IsZero()
}

// IsSet returns true for valid value.
func (e Time) IsSet() bool {
	return e.Valid
}
