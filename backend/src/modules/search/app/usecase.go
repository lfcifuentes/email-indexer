package searchapp

import (
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
	Emails []emaildomain.Email `json:"emails"`
	Total  int                 `json:"total"`
}

func (s SearchUseCases) SearchEmails(query string) (*SearchResponse, error) {
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
		From: 0,
		Size: 20,
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
	// make response
	response := SearchResponse{
		Emails: emails,
		Total:  zincResults.Hits.Total.Value,
	}
	return &response, nil
}
