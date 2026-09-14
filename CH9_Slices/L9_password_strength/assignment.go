package main

func isValidPassword(password string) bool {
	size := len(password)
	isCorrectSize := false
	isContainChar := false
	isContainDigit := false

	if size >= 5 && size <= 12 {
		isCorrectSize = true
	}

	for i := 0; i < size; i++ {
		ch := password[i]
		if ch >= 'A' && ch <= 'Z' {
			isContainChar = true
		}
		if ch >= '0' && ch <= '9' {
			isContainDigit = true
		}
	}

	if isCorrectSize && isContainChar && isContainDigit {
		return true
	}

	return false
}
