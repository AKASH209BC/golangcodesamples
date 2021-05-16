package main

import "fmt"

// type keyword is used for user defined data types
// the underlying type of cash is integer
// x ia variable of type cash
type cash int
var x cash

func main(){
	// %T is used to get the variable data type
	// \n is an escape character which is used to add a new line
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x = 42
	fmt.Println(x)
}
