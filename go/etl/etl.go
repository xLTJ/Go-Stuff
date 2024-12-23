package etl

import "strings"

// Transform changes the date format from one-to-many to one-to-one
func Transform(in map[int][]string) map[string]int {
	newMap := map[string]int{}

	for value, letters := range in {
		for _, letter := range letters {
			newMap[strings.ToLower(letter)] = value
		}
	}

	return newMap
}
