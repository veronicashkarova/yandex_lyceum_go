package main

func FiveSteps(array [5]int) []int {
	reversedSteps := make([]int, 5)
	for i := 0; i < 5; i = i + 1 {
		reversedSteps[i] = array[5-1-i]
	}
	return reversedSteps
}
