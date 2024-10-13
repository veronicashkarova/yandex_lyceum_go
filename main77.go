package main

import (
	"fmt"
)

func main() {
	var x int
	var y int = 0
	fmt.Print()
	fmt.Scanln(&x)
	for i := 1; i <= x; i++ {
		if i%3 == 0 || i%5 == 0 {
			continue

		}

		y = y + i
	}
	fmt.Println(y)
}
