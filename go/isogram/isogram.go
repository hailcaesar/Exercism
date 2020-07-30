package isogram

import "strings"

//IsIsogram determines if input string is an isogram
func IsIsogram(input string) bool {
	input = strings.ToLower(input)
	counter := make(map[rune]int)
	ans := true

	for _, ch := range input {
		if ch != ' ' && ch != '-' && counter[ch] > 0 {
			ans = false
			break
		} else {
			counter[ch]++
		}
	}

	return ans
}
