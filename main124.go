package main

import(
"errors"
)

func Factorial(n int) (int, error) {
	if n < 0 {
		err := errors.New("factorial is not defined for negative numbers")
		return 0, err
}

if n == 0 || n == 1 {
	return 1, nil
}

fact, _ := Factorial(n-1)
return n * fact, nil
}