package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})
	for _, msg := range messages {
		words := strings.Fields(msg)
		for _, word := range words {
			lowerWord := strings.ToLower(word)
			if _, ok := distinctWords[lowerWord]; !ok {
				distinctWords[lowerWord] = struct{}{}
			}
		}
	}

	return len(distinctWords)
}
