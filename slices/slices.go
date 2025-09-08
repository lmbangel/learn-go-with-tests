package main

func SumSlices(s [5]int) int {

	sum := 0
	for _, v := range s {
		sum = sum + v
	}
	return sum
}
