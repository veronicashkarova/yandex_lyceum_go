package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Print()
	fmt.Scanln(&x)
	for i := 1; i <= x; i++ {
		if i%3 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
