package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnvFile reads a .env file and sets the environment variables without relying on external packages.
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
