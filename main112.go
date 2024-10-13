package main

import "fmt"

type Animal interface {
    MakeSound()
}
type Dog struct {

}

func (dog Dog) MakeSound() {
	fmt.Print("Гав")
}

type Сat struct {
	
}

func (cat Сat) MakeSound() {
	fmt.Print("Мяу")
}

