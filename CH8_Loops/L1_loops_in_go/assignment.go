package main

func bulkSend(numMessages int) float64 {
	var result float64
	for num := 0; num < numMessages; num++ {
		result += 1.0 + float64(num)/100
	}
	return result
}
