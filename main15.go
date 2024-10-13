package main

func FiveSteps(array [5]int) [5]int {
	reversed := make([]int, len(array))
	for i := 0; i < len(array); i = i + 1 {
		reversed[i] = array[len(array)-1-i]
	}
	return reversed
}
