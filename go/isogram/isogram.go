//Package isogram provides methods to check isograms
package isogram

import (
	"strings"
	"unicode"
)

//IsIsogram checks is the given text is an isogram
//It returns wheather the given text is an isogram or not
func IsIsogram(value string) bool {
	valueUpper := strings.ToUpper(value)
	chars := map[rune]bool{}
	for _, char := range valueUpper {
		if unicode.IsSpace(char) || unicode.IsPunct(char) || unicode.IsSymbol(char) {
			continue
		}
		if chars[char] {
			return false
		}
		chars[char] = true
	}
	return true
}
