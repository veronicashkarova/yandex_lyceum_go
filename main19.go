package main

func FindMinMaxInArray(array [10]int) (int, int) {
	var Min = array [0]
	var Max = array [0]
	for i := 1; i < 10; i = i + 1 {
		if Min < array [i] {
			Min = array[i]
		}
		if Max > array[i]{
			Max = array[i]
		}
	}
	return Min, Max	
}
