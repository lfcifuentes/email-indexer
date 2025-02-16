package zinc

import (
	"encoding/json"
	"fmt"

	searchdomain "github.com/lfcifuentes/email-indexer/backend/src/modules/search/domain"
	"github.com/lfcifuentes/email-indexer/backend/src/services/api"
)

const Index string = "test_otro"

// ZincService represents a service for interacting with the Zinc API.
// It contains the necessary configuration for making API requests.
//
// Fields:
// - apiURL: The base URL of the Zinc API.
// - apis: An adapter for making API requests.
// - username: The username for authenticating with the Zinc API.
// - password: The password for authenticating with the Zinc API.
type ZincService struct {
	apiURL   string
	apis     api.ApiServiceAdapter
	username string
	password string
}

// NewZincService creates a new instance of ZincService configured with the provided API URL.
func NewZincService(
	apiURL string,
	authUsername string,
	authPassword string,
) *ZincService {
	return &ZincService{
		apiURL: apiURL,
		apis: api.NewApiService(
			authUsername,
			authPassword,
		),
		username: "",
		password: "",
	}
}

// Search sends a search query to the ZincSearch API using the generic ApiService and returns the parsed response.
func (z *ZincService) Search(index string, searchReq searchdomain.SearchRequest) (*searchdomain.SearchResponse, error) {
	// Construct the request URL. Here we assume the search endpoint is "/search".
	url := fmt.Sprintf("%s/es/%s/_search", z.apiURL, index)

	// Prepare headers for the request.
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Execute the request using the ApiService.
	respBody, err := z.apis.PostWithAuth(url, headers, searchReq)
	if err != nil {
		return nil, fmt.Errorf("error sending search request: %w", err)
	}

	// Parse the response.
	var searchResp searchdomain.SearchResponse
	if err = json.Unmarshal(respBody, &searchResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling search response: %w", err)
	}

	return &searchResp, nil
}

// Ping checks if the ZincSearch API is reachable.
func (z *ZincService) Ping() error {
	// Construct the request URL. Here we assume the index endpoint is "/index".
	url := fmt.Sprintf("%s/healthz", z.apiURL)
	_, err := z.apis.Get(url, nil)
	if err != nil {
		return fmt.Errorf("error pinging ZincSearch API: %w", err)
	}
	return nil
}

// Index sends a bulk index request to the ZincSearch API using the generic ApiService and returns the parsed response.
func (z *ZincService) Index(index string, data []map[string]interface{}) (*searchdomain.IndexResponse, error) {
	// Prepare the complete payload with index and records.
	payloadObj := map[string]interface{}{
		"index":   index,
		"records": data,
	}
	// Construct the request URL. Here we assume the index endpoint is "/index".
	url := fmt.Sprintf("%s/api/_bulkv2", z.apiURL)
	// Prepare headers for the request.
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	// Execute the request using the ApiService.
	respBody, err := z.apis.PostWithAuth(url, headers, payloadObj)
	if err != nil {
		return nil, fmt.Errorf("error sending index request: %w", err)
	}

	// Parse the response.
	var searchResp searchdomain.IndexResponse
	if err = json.Unmarshal(respBody, &searchResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling search response: %w", err)
	}

	return &searchResp, nil
}
