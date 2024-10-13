package main

import (
	"strconv"
)

func ConcatStringsAndInt(str1, str2 string, num int) string {
	result := ""
	result += str1
	result += str2
	result += strconv.Itoa(num)

	return result
}
