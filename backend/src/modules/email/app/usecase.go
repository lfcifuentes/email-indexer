package emailapp

import (
	"fmt"
	"sync"

	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

// EmailUseCases represents the email use cases
type EmailUseCases struct {
	EmailServices *email.EmailServices
	ZincServices  *zinc.ZincService
}

// NewUseCase creates a new email use case
// It receives the email services and zinc services
// It returns a new email use case
//
// Parameters:
//   - emailServices: the email services
//   - zincServices: the zinc services
//
// Returns:
//   - a new email use case
func NewUseCase(
	emailServices *email.EmailServices,
	zincServices *zinc.ZincService,
) *EmailUseCases {
	return &EmailUseCases{
		EmailServices: emailServices,
		ZincServices:  zincServices,
	}
}

// RefreshUserEmails refreshes the user emails
// It processes the user emails and sends them to zinc
//
// Parameters:
//   - username: the username
//
// Returns:
//   - the number of emails processed
//   - an error if any
func (u *EmailUseCases) RefreshUserEmails(username string) (int, error) {
	data, err := u.EmailServices.ProcessUserEmails(username)
	if err != nil {
		return 0, err
	}

	// convert data to data []map[string]interface{}
	convertedData := make([]map[string]interface{}, len(data))
	for i, v := range data {
		convertedData[i] = v.ToMap()
		// set the document id
		convertedData[i]["_id"] = fmt.Sprintf("%s-%s", username, v.MessageID)
	}
	// send data to zinc
	response, err := u.ZincServices.Index("test_otro", convertedData)
	if err != nil {
		return 0, err
	}

	return response.RecordCount, nil
}

// RefreshAllEmails refreshes all the emails
// It processes all the emails and sends them to zinc
//
// Returns:
//   - an error if any
func (u *EmailUseCases) RefreshAllEmails() error {
	// Get the list of users (folders)
	users, err := u.EmailServices.GetUserFolders()
	if err != nil {
		return err
	}
	// Use a WaitGroup and a channel to collect errors from goroutines.
	var wg sync.WaitGroup
	errChan := make(chan error, len(users))
	// Limit the number of concurrent goroutines to 10.
	concurrencyLimit := 10
	sem := make(chan struct{}, concurrencyLimit)

	// store the total number of emails imported
	var totalImported int

	for _, user := range users {
		wg.Add(1)
		// Launch a goroutine for each user
		go func(username string) {
			defer wg.Done()
			// start semaphore to limit concurrency
			sem <- struct{}{}
			// release semaphore when done
			defer func() { <-sem }()

			totalRefresn, err := u.RefreshUserEmails(username)
			if err != nil {
				errChan <- err
			}
			totalImported += totalRefresn
		}(user)
	}

	wg.Wait()
	close(errChan)

	// If any error occurred, return the first one.
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

// GetUserNames returns the list of user names found in the mail directory
//
// Returns:
//   - the list of user names
//   - an error if any
func (u *EmailUseCases) GetUserNames() ([]string, error) {
	users, err := u.EmailServices.GetUserFolders()
	if err != nil {
		return nil, err
	}
	return users, nil
}
