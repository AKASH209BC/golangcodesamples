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
	// Create a map which stores the value of person last names with key as last name
	lastname_map:= map[string]person{}
	for x:=0;x<2;x++{
		lastname_map[persons[x].last_name]=persons[x]
	}
	for i,v:=range lastname_map{
		fmt.Println(i,v)
	}
}
