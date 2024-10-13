package main

import (
	"fmt"
	"math"
)

func main() {
	SqRoots()
}
func SqRoots() {
	var a, b, c float64
	fmt.Print()
	fmt.Scanln(&a, &b, &c)
	var x float64 = 0.0
	var y float64 = 0.0
	discriminant := math.Pow(b, 2) - 4.0*a*c
	if discriminant < 0.0 {
		fmt.Println(x, y)
	}
	if discriminant == 0.0 {
		x = -b / (2 * a)
		y = 0.0
		fmt.Println(x)
	}
	if discriminant > 0.0 {
		x = (-b - math.Sqrt(discriminant)) / (2 * a)
		y = (-b + math.Sqrt(discriminant)) / (2 * a)
		fmt.Println(x, y)
	}
}
