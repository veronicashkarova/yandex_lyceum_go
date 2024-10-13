package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(GetUTFLength([]byte {0xFF}))
}

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
