package main

func Join(nums1, nums2 []int) []int {
	slice := make([]int, len(nums1)+len(nums2))
	for i := 0; i < len(nums1); i++ {
		slice[i] = nums1[i]
	}
	for i := 0; i < len(nums2); i++ {
		slice[i+len(nums1)] = nums2[i]
	}
	return slice
}
