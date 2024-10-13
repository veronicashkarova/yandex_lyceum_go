package main

import "fmt"

func main() {
	var score int
	fmt.Print()
	fmt.Scanln(&score)
	switch {
	case score > 0  && score % 2 == 0:
		fmt.Println("Число положительное и четное")
	case score < 0  && score % 2 == 0:
		fmt.Println("Число отрицательное и четное")
	case score > 0  && score % 2 != 0:
		fmt.Println("Число положительное и нечетное")
	case score < 0  && score % 2 != 0:
		fmt.Println("Число отрицательное и нечетное")	
	default:
		fmt.Println("Число равно нулю")
	}	
}
