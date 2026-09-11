package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	result := []float64{}
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			result = append(result, costs[i].value)
		}
	}
	return result
}
