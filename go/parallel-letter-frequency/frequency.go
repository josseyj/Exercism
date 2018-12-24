//Package letter provides methods to calculate letter frequency in a text
package letter

import "sync"

//Letter frequence map type
type FreqMap map[rune]int

//Frequency calculates the letter frequency in a text
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency calculates the letter frequency in the given texts
func ConcurrentFrequency(values []string) FreqMap {
	result := FreqMap{}
	letters := make(chan rune, 10)
	var wg sync.WaitGroup
	for _, value := range values {
		wg.Add(1)
		go freq(value, letters, &wg)
	}
	go func() {
		for letter := range letters {
			result[letter]++
		}
	}()
	wg.Wait()
	close(letters)
	return result
}

func freq(value string, letters chan rune, wg *sync.WaitGroup) {
	for _, letter := range value {
		letters <- letter
	}
	wg.Done()
}
