package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, userName := range messagedUsers {
		if _, ok := validUsers[userName]; ok {
			validUsers[userName]++
		}
	}
}
