package main

func SliceCopy(nums []int) []int {
	slice := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		slice[i] = nums[i]
	}
	return slice
}
