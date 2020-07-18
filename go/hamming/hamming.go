//Package to calculate the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

//Calculate hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("length mismatch:  DNA strands provided are not of same length")
	}

	var distance int
	for idx, _ := range a {
		if a[idx] != b[idx] {
			distance++
		}
	}

	return distance, nil
}
