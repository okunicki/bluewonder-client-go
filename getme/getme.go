package getme

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	Token      = "878ohgj78"
	DefaultURL = "https://mockapigwd.azurewebsites.net/"
)

// MeResponse represents the structure of the response from the API
type MeResponse struct {
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Groups []string `json:"groups"`
}

// GetMe retrieves data from the specified path using the global bearer token and default URL
func GetMe(path string) (*MeResponse, error) {
	// Construct the API URL
	apiURL := DefaultURL + path

	// Create a new HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set the Authorization header with the global bearer token
	req.Header.Set("Authorization", "Bearer "+Token)

	// Perform the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check the HTTP status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal JSON into a struct
	var data MeResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return &data, nil
}
