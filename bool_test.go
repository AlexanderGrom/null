package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_MarshalJSON(t *testing.T) {
	var testCases = []struct {
		value    Bool
		expected []byte
	}{
		{
			value:    BoolFrom(true),
			expected: []byte("true"),
		}, {
			value:    BoolFrom(false),
			expected: []byte("false"),
		}, {
			value:    NewBool(true, true),
			expected: []byte("true"),
		}, {
			value:    NewBool(true, false),
			expected: NullBytes,
		}, {
			value:    NewBool(false, false),
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

func TestBool_Marshal(t *testing.T) {
	var testCases = []struct {
		value    Bool
		expected []byte
	}{
		{
			value:    BoolFrom(true),
			expected: []byte("true"),
		}, {
			value:    BoolFrom(false),
			expected: []byte("false"),
		}, {
			value:    NewBool(true, true),
			expected: []byte("true"),
		}, {
			value:    NewBool(true, false),
			expected: NullBytes,
		}, {
			value:    NewBool(false, false),
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

func TestBool_UnmarshalJSON(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Bool
	}{
		{
			rawData:  []byte("true"),
			expected: NewBool(true, true),
		}, {
			rawData:  []byte("false"),
			expected: NewBool(false, true),
		}, {
			rawData:  []byte("null"),
			expected: NewBool(false, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Bool
			var err = value.UnmarshalJSON(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestBool_UnmarshalJSON_Error(t *testing.T) {
	var testCases = []struct {
		rawData []byte
	}{
		{
			rawData: []byte(""),
		}, {
			rawData: []byte("hello"),
		}, {
			rawData: []byte("12345"),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Bool
			var err = value.UnmarshalJSON(tt.rawData)
			assert.Error(t, err)
		})
	}
}

func TestBool_MarshalText(t *testing.T) {
	var testCases = []struct {
		value    Bool
		expected []byte
	}{
		{
			value:    BoolFrom(true),
			expected: []byte("true"),
		}, {
			value:    BoolFrom(false),
			expected: []byte("false"),
		}, {
			value:    NewBool(true, true),
			expected: []byte("true"),
		}, {
			value:    NewBool(true, false),
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

func TestBool_UnmarshalText(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Bool
	}{
		{
			rawData:  []byte("true"),
			expected: NewBool(true, true),
		}, {
			rawData:  []byte("false"),
			expected: NewBool(false, true),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Bool
			var err = value.UnmarshalText(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}
