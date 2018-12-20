//Package raindrops speaks in rain drops
package raindrops

import "strconv"

//Convert converts the value to raindrop language
func Convert(value int) string {
	var result string
	if value%3 == 0 {
		result += "Pling"
	}
	if value%5 == 0 {
		result += "Plang"
	}
	if value%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(value)
	}
	return result
}
