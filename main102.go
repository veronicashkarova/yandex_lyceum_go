package main

import "fmt"

func main(){
	Employee{"","", 200, 10.0223}.CalculateTotalSalary()
}

type Employee struct {
	name (string)
	position (string)
	salary (float64)
	bonus (float64)
}

func (p Employee )CalculateTotalSalary () {
	fmt.Println("Employee:", p.name)
	fmt.Println("Position:", p.position)
	fmt.Print("Total Salary:")
	fmt.Printf("%.2f", p.salary+p.bonus)
}