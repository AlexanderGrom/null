package null

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTime_Marshal(t *testing.T) {
	var tm = time.Unix(1583353779, 0).UTC()

	var testCases = []struct {
		value    Time
		expected []byte
	}{
		{
			value:    TimeFrom(tm),
			expected: []byte(`"2020-03-04T20:29:39Z"`),
		}, {
			value:    NewTime(tm, true),
			expected: []byte(`"2020-03-04T20:29:39Z"`),
		}, {
			value:    NewTime(tm, false),
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

func TestTime_Unmarshal(t *testing.T) {
	var tm = time.Unix(1583353779, 0).UTC()

	var testCases = []struct {
		rawData  []byte
		expected Time
	}{
		{
			rawData:  []byte(`"2020-03-04T20:29:39Z"`),
			expected: NewTime(tm, true),
		}, {
			rawData:  []byte("null"),
			expected: NewTime(time.Time{}, false),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Time
			var err = json.Unmarshal(tt.rawData, &value)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestTime_Unmarshal_Error(t *testing.T) {
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
			var value Time
			var err = json.Unmarshal(tt.rawData, &value)
			assert.Error(t, err)
		})
	}
}

func TestTime_MarshalText(t *testing.T) {
	var tm = time.Unix(1583353779, 0).UTC()

	var testCases = []struct {
		value    Time
		expected []byte
	}{
		{
			value:    TimeFrom(tm),
			expected: []byte("2020-03-04T20:29:39Z"),
		}, {
			value:    NewTime(tm, true),
			expected: []byte("2020-03-04T20:29:39Z"),
		}, {
			value:    NewTime(tm, false),
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

func TestTime_UnmarshalText(t *testing.T) {
	var tm = time.Unix(1583353779, 0).UTC()

	var testCases = []struct {
		rawData  []byte
		expected Time
	}{
		{
			rawData:  []byte("2020-03-04T20:29:39Z"),
			expected: NewTime(tm, true),
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value Time
			var err = value.UnmarshalText(tt.rawData)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, value)
		})
	}
}

func TestTime_IsZero(t *testing.T) {
	var tm = time.Unix(1583353779, 0).UTC()

	var testCases = []struct {
		value    Time
		expected bool
	}{
		{
			value:    TimeFrom(tm),
			expected: false,
		}, {
			value:    NewTime(tm, true),
			expected: false,
		}, {
			value:    NewTime(tm, false),
			expected: false,
		}, {
			value:    NewTime(time.Time{}, false),
			expected: true,
		},
	}

	for i, tt := range testCases {
		t.Run("case"+strconv.Itoa(i+1), func(t *testing.T) {
			var value = tt.value.IsZero()
			assert.Equal(t, tt.expected, value)
		})
	}
}
