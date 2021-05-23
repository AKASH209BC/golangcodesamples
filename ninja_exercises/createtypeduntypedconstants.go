package main

import "fmt"

// const declared using const keyword
const(
	// a is an untyped constant
	a = 123
    // b is a typed constant
	b int = 23
)
func main(){
	fmt.Println(a,b)
}
