package main

func adder() func(int) int {
	var sum int
	return func(val int) int {
		sum += val
		return sum
	}
}
