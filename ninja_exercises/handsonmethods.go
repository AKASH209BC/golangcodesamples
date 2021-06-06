package main

import "fmt"

type personnel struct{
	first string
	last string
	age int
}
// Method
func (p personnel) speakfoo(){
	fmt.Println("My name is", p.first, p.last, "and my age is", p.age)
}

func main(){
	person1:= personnel{
		first: "Akash",
		last: "Sharma",
		age: 25,
	}
	person1.speakfoo()
}