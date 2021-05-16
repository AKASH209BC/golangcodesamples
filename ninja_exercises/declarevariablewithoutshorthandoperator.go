package main

import "fmt"

// x, y and z have package level scope
// adding type during declaration attaches a type with the variable
// var name would also work except that it does not have a type attached
var x int
var y string
var z bool

func main(){
	// if we don't assign any value to variable then it takes the default value
	// In go programming language default value is also called zero value
	// for int type it is 0, for string it is an empty string and for boolean it is false
	fmt.Println(x,y,z)
	fmt.Println("The value of x is ", x)
	fmt.Println("The value of y is ", y)
	fmt.Println("The value of z is ", z)
	x = 42
	y = "James Bond"
	z = true
    // Using the fmt package
	fmt.Println(x,y,z)
	fmt.Println("The value of x is ", x)
	fmt.Println("The value of y is ", y)
	fmt.Println("The value of z is ", z)
}
