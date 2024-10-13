package main

func SumDigitsRecursive(n int) int {
	if n == 0 {return 0}
	var sum = n % 10
	return sum + SumDigitsRecursive(n/10)
}
