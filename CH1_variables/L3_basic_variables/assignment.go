package main

import "fmt"

func main() {
	var username string
	username = "eddie_cabot"
	var firstRune rune
	firstRune = '3'

	var isAdmin bool
	isAdmin = true

	var permissions int
	permissions = 0x1F

	var costPerSMS float64
	costPerSMS = 0.05

	fmt.Println("username:", username)
	fmt.Println("isAdmin:", isAdmin)
	fmt.Println("permissions:", permissions)
	fmt.Println("costPerSMS:", costPerSMS)
	fmt.Println("firstRune:", firstRune)
}
