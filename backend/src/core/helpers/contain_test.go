package helpers

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		sl       []string
		search   string
		expected bool
	}{
		{"contains element", []string{"apple", "banana", "cherry"}, "banana", true},
		{"does not contain element", []string{"apple", "banana", "cherry"}, "grape", false},
		{"empty slice", []string{}, "apple", false},
		{"single element match", []string{"apple"}, "apple", true},
		{"single element no match", []string{"banana"}, "apple", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.sl, tt.search)
			if result != tt.expected {
				t.Errorf("Contains(%v, %s) = %v; want %v", tt.sl, tt.search, result, tt.expected)
			}
		})
	}
}
