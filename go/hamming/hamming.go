//Package hamming provides utility methods for hamming distance
package hamming

import "errors"

//Distance calculate the hamming distance between two DNA strands
//it retruns the hamming distance and any error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("invalid input - length mismatch")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
