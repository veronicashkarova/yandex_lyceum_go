package main

import ("unicode"
 "fmt"
)

func main() {
	fmt.Print(IsLatin("yandexlyceumgolang"))
}
func IsLatin(input string) bool {
	for _, r := range input {
		if !unicode.Is(unicode.Latin, r) {
			return false
		}
	}
	return true
}
