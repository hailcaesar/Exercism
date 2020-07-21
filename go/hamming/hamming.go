//Package hamming is used to calculate the Hamming Distance between two DNA strands.
package hamming

import (
	"errors"
)

//Distance function used to calculate hamming distance
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)

	if len(ar) != len(br) {
		return -1, errors.New("length mismatch:  DNA strands provided are not of same length")
	}

	var distance int
	for idx := range ar {
		if ar[idx] != br[idx] {
			distance++
		}
	}

	return distance, nil
}
