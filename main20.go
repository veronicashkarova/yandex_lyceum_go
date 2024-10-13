package main

func SumOfArray(array [6]int) int {
	var sum = 0
	for i := 0; i < 6; i = i + 1 {
		sum = sum + array[i]
	}
	return sum 
}