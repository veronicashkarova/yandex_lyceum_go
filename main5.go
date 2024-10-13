package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print()
	_, err := fmt.Scanln(&x, &y, &z)
	if err != nil {
		fmt.Print("Некорректный ввод")
	} else {
		if x == y && y == z && x == z {
			fmt.Print("Все числа равны")
		} else {
			if x == y || y == z || x == z {
				fmt.Print("Два числа равны")
			} else {
				if x != y && y != z && z != x {
					fmt.Print("Все числа разные")
				} else {
					fmt.Print("Некорректный ввод")
				}
			}

		}

	}
}
