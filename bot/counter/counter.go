package counter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Counter struct {
	Champion string   `json:"champion"`
	Counters []string `json:"counters"`
}

func FetchCounter(counterUrl string, champion string) (Counter, error) {
	// path escape worked with dr mundo

	encodedChampion := url.QueryEscape(champion)
	// Construct the final URL with the modified encoded query parameter
	finalURL := counterUrl + "?champion=" + encodedChampion
	resp, err := http.Get(finalURL)
	if err != nil {
		return Counter{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return Counter{}, err
	}

	if resp.StatusCode == 404 {
		fmt.Println(champion + " not found")
		return Counter{}, errors.New(champion + " not found")
	}

	// Parse the JSON response
	var counter Counter
	if err := json.Unmarshal(body, &counter); err != nil {
		fmt.Println("Error parsing JSON:", err)
		responseString := string(body)
		fmt.Println("body is", responseString)
		return Counter{}, err
	}
	return counter, nil

}
