package main

import (
	"fmt"
)

// fibbonachi generator

func outer() func() {

	var step int = 0
	var i, j int = 0, 1

	return func() {
		var n int
		if step < 2 {
			fmt.Println(step)
		} else {
			fmt.Println()
			n = i + j
			fmt.Println(n)
			i = j
			j = n

		}

		step += 1
	}

}

func main() {

	fn := outer()

	fn()
	fn()
	fn()
	fn()
	fn()
	fn()
	fn()
	fn()

}
