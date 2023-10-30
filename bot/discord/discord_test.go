package discord

import (
	"bot/counter"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDaysSinceNow(t *testing.T) {
	tests := []struct {
		name         string
		inputTime    int64
		expectedDays int
	}{
		{"Yesterday", time.Now().Add(-24 * time.Hour).Unix(), 1},
		{"A couple of hours ago", time.Now().Add(-4 * time.Hour).Unix(), 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualDays := daysSinceNow(tc.inputTime)
			assert.Equal(t, tc.expectedDays, actualDays, "Days calculation should match")
		})
	}
}

func TestFormatCounterMessage(t *testing.T) {
	// Create a sample counter object for testing.
	counters := counter.Counter{
		Champion:    "Sample Champion",
		Counters:    []string{"Counter 1", "Counter 2", "Counter 3"},
		LastUpdated: time.Now().Add(-48 * time.Hour).Unix(), // Updated 2 days ago
	}

	expectedMessage := "```\nCounter for Sample Champion:\n1. Counter 1\n2. Counter 2\n3. Counter 3\n\nIt has been 2 days since data was refreshed```"

	// Call the function and use testify's assert library for comparison.
	actualMessage := formatCounterMessage(counters)
	assert.Equal(t, expectedMessage, actualMessage, "Formatted message should match")
}
