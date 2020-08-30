package luhn

import (
	"regexp"
)

//Valid determines if a string meets Luhn's formula requirements
func Valid(input string) bool {
	input = regexp.MustCompile("\\s").ReplaceAllLiteralString(input, "")
	if len(input) < 2 {
		return false
	}
	sum := 0
	odd := len(input) % 2
	for i, ch := range input {
		if ch < '0' || ch > '9' {
			return false
		}
		d := int(ch - '0')
		if i%2 == odd {
			d *= 2
			if d >= 10 {
				d -= 9
			}
		}
		sum += d
	}

	return sum%10 == 0
}
