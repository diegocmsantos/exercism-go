package etl

import "strings"

func Transform(legacy map[int][]string) map[string]int {
	ret := make(map[string]int, len(legacy))
	for value, letters := range legacy {
		for _, letter := range letters {
			ret[strings.ToLower(letter)] = value
		}
	}
	return ret
}