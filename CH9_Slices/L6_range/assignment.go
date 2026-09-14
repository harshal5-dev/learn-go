package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for index, message := range msg {
		for _, badWord := range badWords {
			if badWord == message {
				return index
			}
		}
	}
	return -1
}
