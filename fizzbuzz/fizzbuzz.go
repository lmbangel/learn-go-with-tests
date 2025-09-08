package main

import "fmt"

func FizzBuzz(n int) []string {
	var str []string

	for i := n; i > 0; i-- {
		if i%3 == 0 && i%5 == 0 {
			str = append([]string{"FizzBuzz"}, str...)
		}
		if i%3 == 0 && i%5 != 0 {
			str = append([]string{"Fizz"}, str...)
		}
		if i%3 != 0 && i%5 == 0 {
			str = append([]string{"Buzz"}, str...)
		}
		if i%3 != 0 && i%5 != 0 {
			str = append([]string{fmt.Sprintf("%d", i)}, str...)
		}
	}

	return str
}

func main() {
	fmt.Println(FizzBuzz(15))
}
