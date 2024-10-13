package main

import "fmt"

func main() {
	fmt.Print(Clean([]int{5, 5}, 5))
}

func Clean(nums []int, x int) []int {
	var k = 0
	for i := 0; i < len(nums); i++ {
		for len(nums) > i+k && nums[i+k] == x {
			k = k + 1
		}
		if len(nums) > i+k {
			nums[i] = nums[i+k]
		}
	}
	return nums[:(len(nums) - k)]
}
