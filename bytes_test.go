package null

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes_MarshalJSON(t *testing.T) {
	var testCases = []struct {
		value    Bytes
		expected string
	}{
		{
			value:    BytesFrom([]byte("hello")),
			expected: `"aGVsbG8="`,
		}, {
			value:    NewBytes([]byte("hello"), true),
			expected: `"aGVsbG8="`,
		}, {
			value:    NewBytes([]byte("hello"), false),
			expected: string(NullBytes),
		}, {
			value:    NewBytes([]byte(""), false),
			expected: string(NullBytes),
		}, {
			value:    NewBytes(nil, false),
			expected: string(NullBytes),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value, err = tt.value.MarshalJSON()

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(value))

		})
	}
}

func TestBytes_Marshal(t *testing.T) {
	var testCases = []struct {
		value    Bytes
		expected string
	}{
		{
			value:    BytesFrom([]byte("hello")),
			expected: `"aGVsbG8="`,
		}, {
			value:    NewBytes([]byte("hello"), true),
			expected: `"aGVsbG8="`,
		}, {
			value:    NewBytes([]byte("hello"), false),
			expected: string(NullBytes),
		}, {
			value:    NewBytes([]byte(""), false),
			expected: string(NullBytes),
		}, {
			value:    NewBytes(nil, false),
			expected: string(NullBytes),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value, err = json.Marshal(tt.value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(value))
		})
	}
}

func TestBytes_UnmarshalJSON(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Bytes
	}{
		{
			rawData:  []byte(`"aGVsbG8="`),
			expected: NewBytes([]byte("hello"), true),
		}, {
			rawData:  []byte("null"),
			expected: NewBytes(nil, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Bytes
			var err = value.UnmarshalJSON(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestBytes_Unmarshal(t *testing.T) {
	var testCases = []struct {
		rawData  []byte
		expected Bytes
	}{
		{
			rawData:  []byte(`"aGVsbG8="`),
			expected: NewBytes([]byte("hello"), true),
		}, {
			rawData:  []byte(`""`),
			expected: NewBytes([]byte(""), true),
		}, {
			rawData:  []byte("null"),
			expected: NewBytes(nil, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Bytes
			var err = json.Unmarshal(tt.rawData, &value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestBytes_MarshalUnmarshal(t *testing.T) {
	type s struct {
		B1 []byte
		B2 Bytes
	}

	var e1 = s{
		B1: []byte("hello"),
		B2: BytesFrom([]byte("hello")),
	}

	var value, err = json.Marshal(e1)
	assert.NoError(t, err)

	var e2 s
	err = json.Unmarshal(value, &e2)
	assert.NoError(t, err)

	assert.Equal(t, e1, e2)
}
