package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_MarshalJSON(t *testing.T) {
	var testCases = []struct {
		value    String
		expected []byte
	}{
		{
			value:    StringFrom("hello"),
			expected: []byte(`"hello"`),
		}, {
			value:    NewString("hello", true),
			expected: []byte(`"hello"`),
		}, {
			value:    NewString("hello", false),
			expected: NullBytes,
		}, {
			value:    NewString("", false),
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

func TestString_Marshal(t *testing.T) {
	var testCases = []struct {
		value    String
		expected []byte
	}{
		{
			value:    StringFrom("hello"),
			expected: []byte(`"hello"`),
		}, {
			value:    NewString("hello", true),
			expected: []byte(`"hello"`),
		}, {
			value:    NewString("hello", false),
			expected: NullBytes,
		}, {
			value:    NewString("", false),
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

func TestString_UnmarshalJSON(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected String
	}{
		{
			rawData:  []byte(`"hello"`),
			expected: NewString("hello", true),
		}, {
			rawData:  []byte(`""`),
			expected: NewString("", true),
		}, {
			rawData:  []byte("null"),
			expected: NewString("", false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value String
			var err = value.UnmarshalJSON(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestString_UnmarshalJSON_Error(t *testing.T) {
	var testCases = []struct {
		rawData []byte
	}{
		{
			rawData: []byte(""),
		}, {
			rawData: []byte("0"),
		}, {
			rawData: []byte("true"),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value String
			var err = value.UnmarshalJSON(tt.rawData)
			assert.Error(t, err)
		})
	}
}

func TestString_Unmarshal(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected String
	}{
		{
			rawData:  []byte(`"hello"`),
			expected: NewString("hello", true),
		}, {
			rawData:  []byte(`""`),
			expected: NewString("", true),
		}, {
			rawData:  []byte("null"),
			expected: NewString("", false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value String
			var err = json.Unmarshal(tt.rawData, &value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestString_MarshalText(t *testing.T) {
	var testCases = []struct {
		value    String
		expected []byte
	}{
		{
			value:    StringFrom("hello"),
			expected: []byte("hello"),
		}, {
			value:    NewString("hello", true),
			expected: []byte("hello"),
		}, {
			value:    NewString("hello", false),
			expected: []byte(""),
		}, {
			value:    NewString("", false),
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

func TestString_UnmarshalText(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected String
	}{
		{
			rawData:  []byte("hello"),
			expected: NewString("hello", true),
		}, {
			rawData:  []byte(""),
			expected: NewString("", true),
		}, {
			rawData:  []byte("null"),
			expected: NewString("null", true),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value String
			var err = value.UnmarshalText(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}
