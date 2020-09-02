package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/json"
	"fmt"
)

var (
	_ sql.Scanner              = (*Bytes)(nil)
	_ driver.Valuer            = (*Bytes)(nil)
	_ json.Marshaler           = (*Bytes)(nil)
	_ json.Unmarshaler         = (*Bytes)(nil)
	_ encoding.TextMarshaler   = (*Bytes)(nil)
	_ encoding.TextUnmarshaler = (*Bytes)(nil)
)

// Bytes is a nullable string. Bytes provides SQL and JSON serialization.
type Bytes struct {
	Bytes []byte
	Valid bool
}

// NewBytes creates a new Bytes
func NewBytes(b []byte, v bool) Bytes {
	return Bytes{b, v}
}

// BytesFrom creates a new valid Bytes from []byte.
func BytesFrom(b []byte) Bytes {
	return NewBytes(b, b != nil)
}

// Scan implements sql.Scanner interface.
func (e *Bytes) Scan(value interface{}) error {
	if value == nil {
		e.Bytes, e.Valid = nil, false
		return nil
	}
	switch value := value.(type) {
	case nil:
		e.Bytes = nil
		e.Valid = false
	case []byte:
		// value must be copy
		e.Bytes = make([]byte, len(value))
		copy(e.Bytes, value)
		e.Valid = true
	case string:
		e.Bytes = []byte(value)
		e.Valid = true
	default:
		return fmt.Errorf("cannot convert type %T to bytes", value)
	}
	return nil
}

// Value implements driver.Valuer interface.
func (e Bytes) Value() (driver.Value, error) {
	if !e.Valid {
		return nil, nil
	}
	return e.Bytes, nil
}

// MarshalJSON implements json.Marshaler.
func (e Bytes) MarshalJSON() ([]byte, error) {
	if !e.Valid {
		return NullBytes, nil
	}
	return json.Marshal(e.Bytes)
}

// UnmarshalJSON implements json.Unmarshaler.
func (e *Bytes) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, NullBytes) {
		e.Bytes, e.Valid = nil, false
		return nil
	}

	var s []byte
	if err = json.Unmarshal(data, &s); err != nil {
		return err
	}

	e.Bytes = s
	e.Valid = data != nil
	return err
}

// MarshalText implements encoding.TextMarshaler.
func (e Bytes) MarshalText() ([]byte, error) {
	if !e.Valid {
		return nil, nil
	}
	return e.Bytes, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (e *Bytes) UnmarshalText(text []byte) error {
	e.Bytes = text
	e.Valid = text != nil
	return nil
}

// IsZero returns true for zero value.
func (e Bytes) IsZero() bool {
	return len(e.Bytes) == 0
}

// IsSet returns true for valid value.
func (e Bytes) IsSet() bool {
	return e.Valid
}

// IsSetZero returns true if value is set to zero.
func (e Bytes) IsSetZero() bool {
	return e.IsSet() && e.IsZero()
}

func cloneBytes(b []byte) []byte {
	if b == nil {
		return nil
	}
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
