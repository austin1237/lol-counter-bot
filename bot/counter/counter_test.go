package counter

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchCounter(t *testing.T) {
	// Create a mock HTTP server for testing.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request URL and method.
		if r.URL.Path == "/your_api_endpoint" && r.Method == http.MethodGet {
			// Create a sample Counter response.
			counter := Counter{
				Champion:    "Sample Champion",
				Counters:    []string{"Counter 1", "Counter 2"},
				LastUpdated: 1635548400, // Unix timestamp for a specific date
			}

			// Serialize the Counter to JSON.
			counterJSON, _ := json.Marshal(counter)

			// Simulate a successful HTTP response.
			w.WriteHeader(http.StatusOK)
			w.Write(counterJSON)
		} else {
			// Simulate a 404 response for unknown endpoints.
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Use the mock server's URL as the counterUrl.
	counterUrl := server.URL + "/your_api_endpoint"
	champion := "Sample Champion"

	// Call the function and check the result.
	counter, err := FetchCounter(counterUrl, champion)
	assert.NoError(t, err) // Check that no error occurred
	assert.Equal(t, champion, counter.Champion)
	assert.Len(t, counter.Counters, 2) // Check the length of counters
}
