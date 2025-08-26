package main

import "fmt"

func Adder(v1 int, v2 int) int {
	sum := v1 + v2
	return sum
}

func main() {
	v1 := 6
	v2 := 3

	sum := Adder(v1, v2)
	fmt.Printf("The sum of %d and %d is %d \n", v1, v2, sum)
}
