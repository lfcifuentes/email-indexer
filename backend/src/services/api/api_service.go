package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ApiService provides a generic service to perform API actions (GET, POST, PUT, DELETE)
// by specifying the target URL, request parameters and headers.
type ApiService struct {
	AuthUsername string
	AuthPassword string
}

// NewApiService creates a new instance of ApiService.
func NewApiService(
	authUsername string,
	authPassword string,
) ApiServiceAdapter {
	return ApiService{
		AuthUsername: authUsername,
		AuthPassword: authPassword,
	}
}

// DoRequestWithOptionalAuth performs an HTTP request with the specified method, URL, headers, and body.
// If both username and password are non-empty, it adds Basic Authentication to the request.
// Pass empty strings for username and password if no authentication is required.
//
// Example usage without login:
//
//	apiSvc := NewApiService()
//	headers := map[string]string{
//	    "Content-Type": "application/json",
//	}
//	payload := map[string]interface{}{
//	    "query": "search term",
//	}
//	respBody, err := apiSvc.DoRequestWithOptionalAuth("POST", "http://localhost:4080/search", false, headers, payload)
//	if err != nil {
//	    // handle error
//	}
//
// Example usage with login:
//
//	apiSvc := NewApiService()
//	headers := map[string]string{
//	    "Content-Type": "application/json",
//	}
//	payload := map[string]interface{}{
//	    "query": "search term",
//	}
//	respBody, err := apiSvc.DoRequestWithOptionalAuth("POST", "http://localhost:4080/search", "myUser", "myPass", headers, payload)
//	if err != nil {
//	    // handle error
//	}
func (a ApiService) DoRequestWithOptionalAuth(method, url string, withAuth bool, headers map[string]string, body interface{}) ([]byte, error) {
	var requestBody []byte
	var err error

	// Marshal body to JSON if provided.
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
	}

	// Create the HTTP request.
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// If both username and password are provided, set Basic Auth.
	if withAuth && a.AuthUsername != "" && a.AuthPassword != "" {
		req.SetBasicAuth(a.AuthUsername, a.AuthPassword)
	}

	// Set provided headers.
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	// Default Content-Type header if not provided.
	if _, exists := headers["Content-Type"]; !exists {
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute the HTTP request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing HTTP request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-successful HTTP status codes.
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("non-success status code: %d", resp.StatusCode)
	}

	// Read and return the response body.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return respBody, nil
}

func (a ApiService) Get(url string, headers map[string]string) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("GET", url, false, headers, nil)
}

func (a ApiService) Post(url string, headers map[string]string, body interface{}) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("POST", url, false, headers, body)
}

func (a ApiService) Put(url string, headers map[string]string, body interface{}) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("PUT", url, false, headers, body)
}

func (a ApiService) Delete(url string, headers map[string]string) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("DELETE", url, false, headers, nil)
}

func (a ApiService) GetWithAuth(url string, headers map[string]string) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("GET", url, true, headers, nil)
}

func (a ApiService) PostWithAuth(url string, headers map[string]string, body interface{}) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("POST", url, true, headers, body)
}

func (a ApiService) PutWithAuth(url string, headers map[string]string, body interface{}) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("PUT", url, true, headers, body)
}

func (a ApiService) DeleteWithAuth(url string, headers map[string]string) ([]byte, error) {
	return a.DoRequestWithOptionalAuth("DELETE", url, true, headers, nil)
}
