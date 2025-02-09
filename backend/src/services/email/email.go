package email

import (
	"os"
	"path/filepath"

	coreerrors "github.com/lfcifuentes/email-indexer/backend/src/core/errors"
	emaildomain "github.com/lfcifuentes/email-indexer/backend/src/modules/email/domain"
)

// EmailServices provides functions to work with email files.
type EmailServices struct {
	EmailsPath string
}

// NewEmailServices creates a new instance of EmailServices.
func NewEmailServices(emailsPath string) *EmailServices {
	return &EmailServices{
		EmailsPath: emailsPath,
	}
}

// getUserFolders returns the list of folder names found in the mail directory.
func (es *EmailServices) GetUserFolders() ([]string, error) {
	entries, err := os.ReadDir(es.EmailsPath)
	if err != nil {
		return nil, err
	}

	var users []string
	for _, entry := range entries {
		if entry.IsDir() {
			users = append(users, entry.Name())
		}
	}
	return users, nil
}

// processUserFolder processes all email files in a given user folder and returns a slice of Email.
func (es *EmailServices) processUserFolder(folderPath string) ([]emaildomain.Email, error) {
	var emails []emaildomain.Email

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			email, err := emaildomain.ParseEmailFile(path)
			if err != nil {
				// Skip files that can't be processed.
				return nil
			}
			emails = append(emails, email)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return emails, nil
}

// ProcessEmail processes all email files found in the user folder.
func (es *EmailServices) ProcessUserEmails(username string) ([]emaildomain.Email, error) {
	userFolder := filepath.Join(es.EmailsPath, username)
	// Check if the user folder exists.
	if _, err := os.Stat(userFolder); os.IsNotExist(err) {
		// Return the custom error from coreerrors.
		return nil, coreerrors.NewUsernamePathNotFoundError(err, username)
	}
	emails, err := es.processUserFolder(userFolder)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

// ProcessEmails processes all emails found in the provided directory.
// It returns a consolidated slice of Email from all user folders.
func (es *EmailServices) ProcessAllEmails() ([]emaildomain.Email, error) {
	users, err := es.GetUserFolders()
	if err != nil {
		return nil, err
	}

	var allEmails []emaildomain.Email
	for _, user := range users {
		emails, err := es.ProcessUserEmails(user)
		if err != nil {
			continue
		}
		allEmails = append(allEmails, emails...)
	}
	return allEmails, nil
}
