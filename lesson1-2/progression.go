package main

import (
	"fmt"
)

func p_equation(x int) int {
	return x + 20
}

func p_iterator(start int, equation func(x int) int) func() {

	var p_elem int = start

	return func() {
		var next int
		next = equation(p_elem)

		fmt.Println()
		fmt.Println(next)

		p_elem = next
	}
}

func main() {

	progression_iterator := p_iterator(10, p_equation)

	progression_iterator()
	progression_iterator()
	progression_iterator()
	progression_iterator()
}
