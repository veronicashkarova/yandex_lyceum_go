package main

import "fmt"

func IsPowerOfTwoRecursive(n int) {
	if n == 2 || n == 1{
		fmt.Println("YES")
		return
	}
	if n%2 == 0 {
		IsPowerOfTwoRecursive(n / 2)
	} else {
		if n%2 == 1 {
			fmt.Println("NO")
			return
		}
	}
}
