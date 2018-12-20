//Package scrabble provides the methods for scrabble
package scrabble

import (
	"strings"
)

var scores = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

//Score calculates the scrabble score
//it returns the scrabble score for the given word
func Score(value string) int {
	result := 0
	valueAllUpper := strings.ToUpper(value)
	for i := 0; i < len(valueAllUpper); i++ {
		for key, val := range scores {
			if strings.IndexByte(key, valueAllUpper[i]) >= 0 {
				result += val
			}
		}
	}
	return result
}
