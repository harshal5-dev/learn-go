package main

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		firstLetter := runes[0]
		if _, ok := result[firstLetter]; !ok {
			result[firstLetter] = make(map[string]int)
		}
		result[firstLetter][name]++
	}
	return result
}
