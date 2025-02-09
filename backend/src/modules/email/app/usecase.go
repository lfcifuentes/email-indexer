package emailapp

import (
	"fmt"
	"sync"

	"github.com/lfcifuentes/email-indexer/backend/src/services/email"
	"github.com/lfcifuentes/email-indexer/backend/src/services/zinc"
)

type EmailUseCases struct {
	EmailServices *email.EmailServices
	ZincServices  *zinc.ZincService
}

func NewUseCase(
	emailServices *email.EmailServices,
	zincServices *zinc.ZincService,
) *EmailUseCases {
	return &EmailUseCases{
		EmailServices: emailServices,
		ZincServices:  zincServices,
	}
}

func (u *EmailUseCases) RefreshUserEmails(username string) (int, error) {
	data, err := u.EmailServices.ProcessUserEmails(username)
	if err != nil {
		return 0, err
	}
	// Log start of processing with username and count.
	// slog.Info("Starting email processing", "event", "start_emails_processed", "username", username, "email_count", len(data))

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

	// Log finish of processing with username and count.
	// slog.Info("Completed email processing", "event", "end_emails_processed", "username", username, "email_count", len(data))

	return response.RecordCount, nil
}

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

func (u *EmailUseCases) GetUserNames() ([]string, error) {
	users, err := u.EmailServices.GetUserFolders()
	if err != nil {
		return nil, err
	}
	return users, nil
}
