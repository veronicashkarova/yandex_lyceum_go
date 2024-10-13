package main

import "math"

type Shape interface {
    Area() float64
}
func CalculateArea(s Shape) float64 {
    return s.Area()


func (c *Circle) Area() float64 {
	return math.Pi * c.radius*c.radius
}
type Rectangle struct {
    width  float64
    height float64
    
}


    return r.width * r.height
