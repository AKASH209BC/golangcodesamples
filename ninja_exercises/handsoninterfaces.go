package main

import "fmt"

const PI = 3.14

type circle struct{
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64{
	return PI * c.radius* c.radius
}

func (s square) area() float64{
	return s.side*s.side
}

type shape interface {
	area() float64
}

func info(s shape){
	fmt.Printf(" The area is %v\n", s.area())
}
func main(){
	s1:=square{
		side: 4,
	}
	c1:=circle{
		radius: 2.4,
	}
	info(s1)
	info(c1)
}
