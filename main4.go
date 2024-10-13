package main

import "fmt"

func main() {
	var x, y int
	fmt.Print()
	fmt.Scanln(&x, &y)
	if x > 0 && y > 0 {
		fmt.Print("Оба числа положительные")
	} else {
		if x < 0 && y < 0 {
			fmt.Print("Оба числа отрицательные")
		} else {
			if (x > 0 && y < 0) || (x < 0 && y > 0){
				fmt.Print("Одно число положительное, а другое отрицательное")
			} else {
				if x == 0 || y == 0 {
					fmt.Print("Одно из чисел равно нулю")
				}
			}
		}
	}
}
