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
	_ sql.Scanner              = (*Int64)(nil)
	_ driver.Valuer            = (*Int64)(nil)
	_ json.Marshaler           = (*Int64)(nil)
	_ json.Unmarshaler         = (*Int64)(nil)
	_ encoding.TextMarshaler   = (*Int64)(nil)
	_ encoding.TextUnmarshaler = (*Int64)(nil)
)

// Int64 is a nullable int64. Int64 provides SQL and JSON serialization.
type Int64 struct {
	sql.NullInt64
}

// NewInt64 creates a new Int64
func NewInt64(i int64, v bool) Int64 {
	return Int64{
		sql.NullInt64{i, v},
	}
}

// Int64From creates a new valid Int64 from raw int64.
func Int64From(i int64) Int64 {
	return NewInt64(i, true)
}

// MarshalJSON implements json.Marshaler.
func (e Int64) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.Int64)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Int64) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.Int64, e.Valid = 0, false
		return nil
	}
	err = json.Unmarshal(data, &e.Int64)
	e.Valid = err == nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e Int64) MarshalText() ([]byte, error) {
	if !e.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(e.Int64, 10)), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Int64) UnmarshalText(text []byte) (err error) {
	e.Int64, err = strconv.ParseInt(string(text), 10, 64)
	e.Valid = err == nil
	return nil
}

// IsZero returns true for zero value.
func (e Int64) IsZero() bool {
	return e.Int64 == 0
}

// IsSet returns true for valid value.
func (e Int64) IsSet() bool {
	return e.Valid
}

// IsSetZero returns true if value is set to zero.
func (e Int64) IsSetZero() bool {
	return e.IsSet() && e.IsZero()
}
