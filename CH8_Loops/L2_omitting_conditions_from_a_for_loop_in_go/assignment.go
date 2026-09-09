package main

func maxMessages(thresh int) int {
	var total int
	for msg := 0; ; msg++ {
		total += 100 + msg
		if total > thresh {
			return msg
		}
	}
}
