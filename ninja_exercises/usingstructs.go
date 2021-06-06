package main

import "fmt"
// Create a struct of type person with fields first_name , last_name and icecream_flavors
// icecream_flavors is of type slice type
type person struct {
	first_name string
	last_name string
	icecream_flavors []string
}
func main(){
	person1:= person{
		first_name: "Akash",
		last_name: "Sharma",
		icecream_flavors: []string{"vanilla"},
	}
	person2:= person{
		first_name: "Punit",
		last_name: "Pathak",
		icecream_flavors: []string{"chocolate"},
	}
	persons:= [2]person{person1,person2}
	fmt.Println("Printing Persons values")
	for x:=0;x<2;x++{
		fmt.Println(persons[x])
	}
}
