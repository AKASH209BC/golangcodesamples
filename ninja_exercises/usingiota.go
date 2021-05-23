package main

import "fmt"

const(
	// value of iota is 0 here
	a = 2021 + iota
	// value of iota becomes 1
	b = 2021 + iota
	// value of iota becomes 2
	c = 2021 + iota
	// value of iota becomes 3
	d = 2021 + iota
)

func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
