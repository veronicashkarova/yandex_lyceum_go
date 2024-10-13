package main

import "fmt"

func main(){
	fmt.Print(Mix([]int{0, 1}))
}

func Mix(nums []int) []int {
	slice := make([]int, len(nums))
	var n = len(nums)/2
	for i := 0; i < n; i++{
		slice[i*2] = nums[i]
		slice[i*2+1] = nums[i+n]
	}
	return slice
}
