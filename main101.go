package main

import "fmt"

type Person struct {
	name string
	age int
	address string
}

func main(){

	 m :=Person{"aa",5,"4"}
	 m.Print()
}
func (p Person) Print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}