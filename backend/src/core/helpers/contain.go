package helpers

// Contains
// checks if a slice of strings contains a specific string.
//
// Parameters:
//   - sl: A slice of strings to search within.
//   - name: The string to search for in the slice.
//
// Returns:
//   - bool: Returns true if the string is found in the slice, otherwise false.
func Contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}
