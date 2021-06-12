package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func changeMe(p *person){
	// Two ways to set
	(*p).first="Simon"
	(*p).last="Katich"
	(*p).age=34
	fmt.Printf("The person is %v %v and age is %v\n", p.first,p.last,p.age)
	p.first="Ashley"
	p.last="Giles"
	p.age=12
}

func main(){
	p:=person{
		first:"Akash",
		last:"Sharma",
		age:25,
	}
	fmt.Printf("The person is %v %v and age is %v\n", p.first,p.last,p.age)
	changeMe(&p)
	fmt.Printf("The person is %v %v and age is %v\n", p.first,p.last,p.age)
}