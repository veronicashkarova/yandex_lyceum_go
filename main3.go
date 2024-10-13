package main

import "fmt"

func main() {
	var x,y int
	fmt.Print()
	fmt.Scanln(&x,&y)
	if x > y {
		fmt.Println("Первое число больше второго")
	} else {
		if y > x {
			fmt.Println("Второе число больше первого")
		} else {
			fmt.Println("Числа равны")
		}
	}
}
