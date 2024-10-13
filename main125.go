package main

import (
	"errors"
	"fmt"
)

func IntToBinary(num int) (string, error) {
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	binary := ""

	for num != 0 {
		remainder := num % 2
		binary = fmt.Sprintf("%d%s", remainder, binary)
		num = num / 2
	}
	return binary, nil
}
