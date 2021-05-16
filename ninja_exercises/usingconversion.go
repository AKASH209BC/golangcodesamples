package main

import "fmt"

// type keyword is used for user defined data types
// the underlying type of cash is integer
// x ia variable of type cash
type cash int
var x cash
var y int

func main(){
	// %T is used to get the variable data type
	// \n is an escape character which is used to add a new line
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
	// converting the value of x which is of type cash to int
	// go does not do cast but it only does conversion
	y = int(x)
	fmt.Printf("%T\n",y)
	fmt.Println(y)
	// the below x datatype as cash because go is a statically typed language
	// statically typed means the data types don't change after assigning
	// only the value of x did get converted to int not x's datatype
	fmt.Printf("%T\n",x)

}

