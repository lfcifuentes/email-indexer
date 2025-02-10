package searchapp

import (
	"strconv"

	emaildomain "github.com/lfcifuentes/email-indexer/backend/src/modules/email/domain"
	searchdomain "github.com/lfcifuentes/email-indexer/backend/src/modules/search/domain"
	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

type SearchUseCases struct {
	EmailServices *email.EmailServices
	ZincServices  *zinc.ZincService
}

func NewUseCase(
	emailServices *email.EmailServices,
	zincServices *zinc.ZincService,
) *SearchUseCases {
	return &SearchUseCases{
		EmailServices: emailServices,
		ZincServices:  zincServices,
	}
}

type SearchResponse struct {
	Emails   []emaildomain.Email `json:"emails"`
	Total    int                 `json:"total"`
	LastPage int                 `json:"last_page"`
}

func (s SearchUseCases) SearchEmails(query string, page string) (*SearchResponse, error) {
	limit := 20
	// calculate page start
	pageStart := 0
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			pageInt = 1
		}
		pageStart = (pageInt - 1) * limit
	}

	request := searchdomain.SearchRequest{
		SearchType: "matchphrase",
		Query: searchdomain.BoolQuery{
			Bool: searchdomain.MustQuery{
				Must: []searchdomain.QueryStringWrapper{
					{
						QueryString: searchdomain.QueryString{
							Query: query,
						},
					},
				},
			},
		},
		From: pageStart,
		Size: limit,
	}
	zincResults, err := s.ZincServices.Search(zinc.Index, request)
	if err != nil {
		return nil, err
	}
	emails := []emaildomain.Email{}
	for _, hit := range zincResults.Hits.Hits {
		source := hit.Source
		tmpEmail := emaildomain.Email{}
		tmpEmail.FromMap(source)
		emails = append(emails, tmpEmail)
	}
	// calculate last page
	lastPage := (zincResults.Hits.Total.Value + limit - 1) / limit
	// make response
	response := SearchResponse{
		Emails:   emails,
		Total:    zincResults.Hits.Total.Value,
		LastPage: lastPage,
	}
	return &response, nil
}
