package main

import "fmt"

func main() {
	var x int
	fmt.Print()
	fmt.Scanln(&x)
	for i := 1; i <= 10; i++ {
		var y = x * i
		fmt.Println(x, " * ", i, " = ", y)
	}
}
