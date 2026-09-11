package main

func getMessageCosts(messages []string) []float64 {
	messagesCost := make([]float64, len(messages))
	for i := 0; i < len(messagesCost); i++ {
		messagesCost[i] = float64(len(messages[i])) * 0.01
	}
	return messagesCost
}
