//Package twelve
package twelve

import "fmt"

var gifts = [12]string{
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

var countStr = [12]string{
	"a",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var indexStr = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

const lyricsTemplate string = "On the %s day of Christmas my true love gave to me: %s."

func Song() string {
	var result string
	for i := 1; i <= 12; i++ {
		result += Verse(i)
		result += "\n"
	}
	return result
}

func Verse(day int) string {
	return fmt.Sprintf(lyricsTemplate, indexStr[day-1], giftsForTheDay(day))
}

func giftsForTheDay(day int) string {
	var result string
	for i := day; i > 0; i-- {
		result += fmt.Sprintf("%s %s", countStr[i-1], gifts[i-1])
		if i > 1 {
			result += ", "
			if i == 2 {
				result += "and "
			}
		}
	}
	return result
}
