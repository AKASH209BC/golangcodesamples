package main

import "fmt"

// x, y and z have package level scope
var x = 42
var y = "James Bond"
var z = true

func main(){
	// Sprintf can be used to return a formatted String
	// %v gives the value of the variable
	// \t is an escape character used for giving tab space
	s:= fmt.Sprintf("%v\t%v\t%v\t",x,y,z)
	fmt.Println(s)
}
