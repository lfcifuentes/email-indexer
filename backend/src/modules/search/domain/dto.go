package searchdomain

// SearchRequest represents the payload sent to the ZincSearch API for searching.
type SearchRequest struct {
	SearchType string    `json:"search_type"`
	Query      BoolQuery `json:"query"`
	From       int       `json:"from"`
	Size       int       `json:"size"`
}
type SearchRequestQuery struct {
	Query BoolQuery `json:"query"`
}

type BoolQuery struct {
	Bool MustQuery `json:"bool"`
}

type MustQuery struct {
	Must []QueryStringWrapper `json:"must"`
}

type QueryStringWrapper struct {
	QueryString QueryString `json:"query_string"`
}

type QueryString struct {
	Query string `json:"query"`
}

// SearchResponse represents the response received from the ZincSearch API.
// TotalWrapper defines the structure for the total hits in the search response.
type TotalWrapper struct {
	Value int `json:"value"`
}

// Hit represents each individual search result hit.
type Hit struct {
	Index     string                 `json:"_index"`
	Type      string                 `json:"_type"`
	ID        string                 `json:"_id"`
	Score     float64                `json:"_score"`
	Timestamp string                 `json:"@timestamp"`
	Source    map[string]interface{} `json:"_source"`
}

// HitsWrapper defines the structure for the hits object in the search response.
type HitsWrapper struct {
	Total    TotalWrapper `json:"total"`
	MaxScore float64      `json:"max_score"`
	Hits     []Hit        `json:"hits"`
}

// ShardsInfo defines the structure for the shards information in the search response.
type ShardsInfo struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

// SearchResponse represents the response received from the ZincSearch API.
type SearchResponse struct {
	Took     int         `json:"took"`
	TimedOut bool        `json:"timed_out"`
	Shards   ShardsInfo  `json:"_shards"`
	Hits     HitsWrapper `json:"hits"`
}

// SearchResponse represents the response received from the ZincSearch API.
type IndexResponse struct {
	Message     string `json:"message"`
	RecordCount int    `json:"record_count"`
}
