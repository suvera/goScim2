package scim2

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected time.Time
	}{
		{"2006-01-02T15:04:05.000-07:00Z", time.Date(2006, 1, 2, 15, 4, 5, 0, time.FixedZone("MST", -7*3600))},
		{"2006-01-02T15:04:05.000Z", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
		{"2006-01-02T15:04:05Z", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
		{"2006-01-02 15:04:05.000-07:00", time.Date(2006, 1, 2, 15, 4, 5, 0, time.FixedZone("MST", -7*3600))},
		{"2006-01-02 15:04:05.000", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
		{"2006-01-02 15:04:05", time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC)},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, _ := ParseDate(tc.input)
			if !result.Equal(tc.expected) {
				t.Errorf("ParseDate(%s) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestIsDate(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"2006-01-02T15:04:05.000-07:00Z", true},
		{"2006-01-02T15:04:05.000Z", true},
		{"2006-01-02T15:04:05Z", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02T15:04:05.000-07:00Z", true},
		{"2006-01-02T15:04:05.000Z", true},
		{"2006-01-02T15:04:05Z", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02 15:04:05.000-07:00", true},
		{"2006-01-02 15:04:05.000", true},
		{"2006-01-02 15:04:05", true},
		{"2006-01-02T15:04:05.000-07:00Z", true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := IsDate(tc.input)
			if result != tc.expected {
				t.Errorf("IsDate(%s) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
