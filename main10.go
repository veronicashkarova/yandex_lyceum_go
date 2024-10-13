package main

import "fmt"

func main() {
	var x int
	fmt.Print()
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		fmt.Println("Некорректный ввод")
		return
	}

	factorial := 1
	for i := 1; i <= x; i++ {
		factorial *= 1
	}

	fmt.Println(factorial)
}
