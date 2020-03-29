package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64_MarshalJSON(t *testing.T) {
	var testCases = []struct {
		value    Int64
		expected []byte
	}{
		{
			value:    Int64From(777),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, true),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, false),
			expected: NullBytes,
		}, {
			value:    NewInt64(0, false),
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

func TestInt64_Marshal(t *testing.T) {
	var testCases = []struct {
		value    Int64
		expected []byte
	}{
		{
			value:    Int64From(777),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, true),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, false),
			expected: NullBytes,
		}, {
			value:    NewInt64(0, false),
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

func TestInt64_UnmarshalJSON(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Int64
	}{
		{
			rawData:  []byte("777"),
			expected: NewInt64(777, true),
		}, {
			rawData:  []byte("0"),
			expected: NewInt64(0, true),
		}, {
			rawData:  []byte("null"),
			expected: NewInt64(0, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Int64
			var err = value.UnmarshalJSON(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestInt64_UnmarshalJSON_Error(t *testing.T) {
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
			var value Int64
			var err = value.UnmarshalJSON(tt.rawData)
			assert.Error(t, err)
		})
	}
}

func TestInt64_MarshalText(t *testing.T) {
	var testCases = []struct {
		value    Int64
		expected []byte
	}{
		{
			value:    Int64From(777),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, true),
			expected: []byte("777"),
		}, {
			value:    NewInt64(777, false),
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

func TestInt64_UnmarshalText(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Int64
	}{
		{
			rawData:  []byte("777"),
			expected: NewInt64(777, true),
		}, {
			rawData:  []byte("0"),
			expected: NewInt64(0, true),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Int64
			var err = value.UnmarshalText(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}
