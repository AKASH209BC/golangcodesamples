package main

import (
	"fmt"
)

type person struct{
	FirstName string `json:"FirstName"`
	LastName string  `json:"LastName"`
	Age int		`json:"Age"`
}

func (p *person)speak(){
	fmt.Printf("Hello! My name is %v %v and I am %d years old\n",(*p).FirstName,(*p).LastName,(*p).Age)
}

type human interface {
	speak()
}

func saySomething(h human){
	h.speak()
}

func main(){
	fmt.Println("Method sets")
	person1:=person{"Alissa","Porter",23}
	person1.speak()
	saySomething(&person1)
	// can't pass a non-pointer type to a method set with pointer type.
	// Uncommenting the below line will cause error
	// saySomething(person1)
}
