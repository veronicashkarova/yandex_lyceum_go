package main

import "fmt"

func main() {
	var score int
	fmt.Print("Введите число")
	fmt.Scanln(&score)
	if score > 0 {
		fmt.Println("Число положительное")
	} else {
		if score < 0 {
			fmt.Println("Число отрицательное")

		} else {
			fmt.Println("Введен ноль")
		}
	}
}
