package main

import "fmt"

func main() {
	var x int
	var y int = 0
	fmt.Print()
	fmt.Scanln(&x)
	for i := 1; i <= x; i++ {
		y = y + i
	}
	fmt.Println(y)
}
