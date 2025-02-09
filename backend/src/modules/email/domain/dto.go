package emaildomain

import (
	"bufio"
	"os"
	"strings"
)

// Email represents the structure of an email.
type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
}

func (e Email) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"message_id": e.MessageID,
		"date":       e.Date,
		"from":       e.From,
		"to":         e.To,
		"subject":    e.Subject,
		"body":       e.Body,
	}
}
func safeString(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	}
	return ""
}

func (e *Email) FromMap(m map[string]interface{}) {
	e.MessageID = safeString(m["message_id"])
	e.Date = safeString(m["date"])
	e.From = safeString(m["from"])
	e.To = safeString(m["to"])
	e.Subject = safeString(m["subject"])
	e.Body = safeString(m["body"])
}

// parseHeaderLine processes a single header line and updates the corresponding Email field.
func parseHeaderLine(email *Email, line string) {
	switch {
	case strings.HasPrefix(line, "Message-ID: "):
		email.MessageID = strings.TrimPrefix(line, "Message-ID: ")
	case strings.HasPrefix(line, "Date: "):
		email.Date = strings.TrimPrefix(line, "Date: ")
	case strings.HasPrefix(line, "From: "):
		email.From = strings.TrimPrefix(line, "From: ")
	case strings.HasPrefix(line, "To: "):
		email.To = strings.TrimPrefix(line, "To: ")
	case strings.HasPrefix(line, "Subject: "):
		email.Subject = strings.TrimPrefix(line, "Subject: ")
	}
}

// ParseEmailFile opens a file, reads its content line by line,
// and constructs an Email struct by parsing header and body lines.
func ParseEmailFile(path string) (Email, error) {
	var email Email
	var bodyLines []string
	headersEnded := false

	file, err := os.Open(path)
	if err != nil {
		return email, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Detect the end of headers when an empty line is encountered.
		if !headersEnded && strings.TrimSpace(line) == "" {
			headersEnded = true
			continue
		}

		if headersEnded {
			bodyLines = append(bodyLines, line)
		} else {
			parseHeaderLine(&email, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return email, err
	}

	email.Body = strings.Join(bodyLines, "\n")
	return email, nil
}
