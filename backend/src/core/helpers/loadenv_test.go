package helpers

import (
	"os"
	"testing"
)

func TestLoadEnvFile(t *testing.T) {
	// Create a temporary .env file for testing
	envContent := `
		# This is a comment
		FOO=bar
		BAZ="qux"
		INVALID_LINE
		QUOTED='quoted_value'
		API_KEY="SDOjasd8324*234234=="
	`
	tmpfile, err := os.CreateTemp("", "testenv")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(envContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Load the .env file
	err = LoadEnvFile(tmpfile.Name())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Check if the environment variables are set correctly
	tests := []struct {
		key      string
		expected string
	}{
		{"FOO", "bar"},
		{"BAZ", "qux"},
		{"QUOTED", "quoted_value"},
		{"API_KEY", "SDOjasd8324*234234=="},
	}

	for _, tt := range tests {
		if value := os.Getenv(tt.key); value != tt.expected {
			t.Errorf("Expected %s to be %s, got %s", tt.key, tt.expected, value)
		}
	}

	// Check if invalid lines are ignored
	if value := os.Getenv("INVALID_LINE"); value != "" {
		t.Errorf("Expected INVALID_LINE to be empty, got %s", value)
	}
}
