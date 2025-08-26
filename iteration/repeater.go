package main

import (
	"strings"
)

func Repeater(r string, i int) string {
	repeat := ""
	for i > 0 {
		repeat = repeat + r
		i--
	}
	return repeat
}

func Repeater2(r string, i int) string {
	var str strings.Builder
	for i > 0 {
		str.WriteString(r)
		i--
	}
	return str.String()
}
