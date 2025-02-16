package searchdomain

import (
	emaildomain "github.com/lfcifuentes/email-indexer/backend/src/modules/email/domain"
)

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

type SearchEmailResponse struct {
	Emails   []emaildomain.Email `json:"emails"`
	Total    int                 `json:"total"`
	LastPage int                 `json:"last_page"`
}
