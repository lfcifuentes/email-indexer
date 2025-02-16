package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnvFile loads environment variables from a specified file.
// The file should contain key-value pairs in the format "KEY=VALUE".
// Lines that are empty or start with a '#' character are ignored as comments.
// Malformed lines that do not contain an '=' character are logged and skipped.
//
// Parameters:
//   - filename: The path to the file containing the environment variables.
//
// Returns:
//   - error: An error if the file cannot be opened or read, or nil if successful.
func LoadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value (separated by '=')
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Printf("Malformed variable: %s", line)
			continue
		}

		key := strings.TrimSpace(parts[0])
		// Remove surrounding quotes if present
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)

		os.Setenv(key, value)
	}

	return scanner.Err()
}
