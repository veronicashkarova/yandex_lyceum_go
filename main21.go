package main

import "fmt"

func PrettyArrayOutput(array [9]string) {
	for i := 0; i < 9; i = i + 1 {
		var done = ""
		if i < 7 {
			done = " я уже сделал: "
		} else {
			done = " не успел сделать: "
		}

		fmt.Println(i+1, done, array[i])
	}
}
