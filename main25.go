package main

import "errors"
import "fmt"

func main(){
	fmt.Print(UnderLimit([]int{},-1,5))
}

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	slice := make([]int, 0)
	var err error = errors.New("error")
	if nums == nil || len(nums) == 0 {
		return nil, err
	}
	for i := 0; i < len(nums);  i++ {
		if nums[i]<limit {
			slice = append(slice, nums[i])
			if len(slice)==n {
				return slice,nil
			}
		}
	}
	return slice,nil
}
