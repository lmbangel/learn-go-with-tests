package main

import "fmt"

const helloPrefix = "Hello,"

func SayHello(s string,) string {
	if s == "" {
		s = "World"
	}
	return fmt.Sprintf("%s %s", helloPrefix, s)
}

func main() {
	var name string
	fmt.Scanf("%s %s", &name)
	fmt.Println(SayHello(name))
}
