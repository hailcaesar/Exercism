package isogram

import "strings"

//IsIsogram determines if input string is an isogram
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	seen := make(map[rune]bool)

	for _, ch := range input {
		if seen[ch] && ch != ' ' && ch != '-' {
			return false
		}

		seen[ch] = true
	}

	return true
}
