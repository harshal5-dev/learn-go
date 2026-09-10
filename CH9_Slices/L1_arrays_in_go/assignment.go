package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msgs := [3]string{primary, secondary, tertiary}
	primaryLen, secondaryLen, tertiaryLen := len(primary), len(secondary), len(tertiary)
	mesgLents := [3]int{primaryLen, primaryLen + secondaryLen, primaryLen + secondaryLen + tertiaryLen}
	return msgs, mesgLents
}
