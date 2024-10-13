package main

import "fmt"

func main() {
	var score int
	fmt.Print()
	fmt.Scanln(&score)
	if score < 0 {
		fmt.Print("Некорректный ввод")
	} else {
		if score < 10 {
			fmt.Print("Число меньше 10")
		} else {
			if score < 100 {
				fmt.Print("Число меньше 100")
			} else {
				if score < 1000 {
					fmt.Print("Число меньше 1000")
				} else {
					if score > 1000 || score == 1000 {
						fmt.Print("Число больше или равно 1000")
					}

				}
			}
		}
	}
}
