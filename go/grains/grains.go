//Package grains
package grains

import "errors"

//Square returns the value at the given square
//returns an error if the given square is invalid
func Square(value int) (uint64, error) {
	if value <= 0 || value > 64 {
		return 0, errors.New("invalid square index - must be between 1 and 64 (inclusive)")
	}

	result := uint64(1)
	for i := 2; i <= value; i++ {
		result *= 2
	}
	return result, nil
}

//Total calculates the total of all the values in the chess board
func Total() uint64 {
	result := uint64(0)
	for i := 1; i <= 64; i++ {
		valueAtSqaure, _ := Square(i)
		result += valueAtSqaure
	}
	return result
}

//TotalWithRoutines calculates the total of all the values in the chess board
//using go routines
func TotalWithRoutines() uint64 {
	valuesAtSqaure := make(chan uint64, 64)
	for i := 1; i <= 64; i++ {
		go func(v int) {
			valueAtSqaure, _ := Square(v)
			valuesAtSqaure <- valueAtSqaure
		}(i)
	}

	result := uint64(0)
	for i := 1; i <= 64; i++ {
		valueAtSqaure := <-valuesAtSqaure
		result += valueAtSqaure
	}
	return result
}
