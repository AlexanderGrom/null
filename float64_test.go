package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloat64_MarshalJSON(t *testing.T) {
	var testCases = []struct {
		value    Float64
		expected []byte
	}{
		{
			value:    Float64From(0.14),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, true),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, false),
			expected: NullBytes,
		}, {
			value:    NewFloat64(0, false),
			expected: NullBytes,
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value, err = tt.value.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestFloat64_Marshal(t *testing.T) {
	var testCases = []struct {
		value    Float64
		expected []byte
	}{
		{
			value:    Float64From(0.14),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, true),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, false),
			expected: NullBytes,
		}, {
			value:    NewFloat64(0, false),
			expected: NullBytes,
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value, err = json.Marshal(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestFloat64_UnmarshalJSON(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Float64
	}{
		{
			rawData:  []byte("0.14"),
			expected: NewFloat64(0.14, true),
		}, {
			rawData:  []byte("0"),
			expected: NewFloat64(0, true),
		}, {
			rawData:  []byte("null"),
			expected: NewFloat64(0, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Float64
			var err = value.UnmarshalJSON(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestFloat64_UnmarshalJSON_Error(t *testing.T) {
	var testCases = []struct {
		rawData []byte
	}{
		{
			rawData: []byte(""),
		}, {
			rawData: []byte("hello"),
		}, {
			rawData: []byte("true"),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Float64
			var err = value.UnmarshalJSON(tt.rawData)
			assert.Error(t, err)
		})
	}
}

func TestFloat64_MarshalText(t *testing.T) {
	var testCases = []struct {
		value    Float64
		expected []byte
	}{
		{
			value:    Float64From(0.14),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, true),
			expected: []byte("0.14"),
		}, {
			value:    NewFloat64(0.14, false),
			expected: []byte(""),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value, err = tt.value.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestFloat64_UnmarshalText(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Float64
	}{
		{
			rawData:  []byte("0.14"),
			expected: NewFloat64(0.14, true),
		}, {
			rawData:  []byte("0"),
			expected: NewFloat64(0, true),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Float64
			var err = value.UnmarshalText(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}
