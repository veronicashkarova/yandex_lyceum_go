package main

func FiveSteps(array [5]int) [5]int {
	FiveSteps := make([]int, len(array))
	for i := 0; i < 5; i = i + 1 {
		FiveSteps[i] = array[len(array)-1-i]
	}
	return FiveSteps
}
