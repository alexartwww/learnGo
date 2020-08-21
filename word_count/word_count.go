package word_count

import "strings"

func WordCount(s string) map[string]int {
	var result map[string]int
	result = make(map[string]int)
	words := strings.Fields(s)
	for _, value := range words {
		_, ok := result[value]
		if ok {
			result[value]++
		} else {
			result[value] = 1
		}
	}
	return result
}
