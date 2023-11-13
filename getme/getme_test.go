package getme_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/client/getme"
)

func TestGetMe(t *testing.T) {
	// Store the original values
	originalToken := getme.Token
	originalURL := getme.DefaultURL

	// Mocking the HTTP server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check the request
		if r.Method != "GET" {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer "+getme.Token {
			t.Errorf("unexpected Authorization header value")
		}
		if r.URL.String() != originalURL+"/me" {
			t.Errorf("unexpected URL, expected %s, got %s", originalURL+"/me", r.URL.String())
		}

		// Respond with the provided JSON
		responseData := getme.MeResponse{Name: "12345@jdadelivers.com", Type: "service_account", Groups: []string{"group1", "group2"}}
		responseJSON, _ := json.Marshal(responseData)
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	}))
	defer mockServer.Close()

	// Set the default URL and token to the original values
	getme.DefaultURL = originalURL
	getme.Token = originalToken

	// Call the function
	response, err := getme.GetMe("me")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check the response
	expected := &getme.MeResponse{Name: "12345@jdadelivers.com", Type: "service_account", Groups: []string{"group1", "group2"}}
	if response.Name != expected.Name {
		t.Errorf("expected Name %s, got %s", expected.Name, response.Name)
	}
	if response.Type != expected.Type {
		t.Errorf("expected Type %s, got %s", expected.Type, response.Type)
	}
	if len(response.Groups) != len(expected.Groups) {
		t.Errorf("expected %d groups, got %d", len(expected.Groups), len(response.Groups))
	}
	for i := range expected.Groups {
		if response.Groups[i] != expected.Groups[i] {
			t.Errorf("expected group %s, got %s", expected.Groups[i], response.Groups[i])
		}
	}
}
