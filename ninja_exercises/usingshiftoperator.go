package main

import "fmt"

func main(){
	a:=12
	// using shift operator
	b:= a<<1
	fmt.Printf("%d\t%b\t%#x\n",a,a,a)
	fmt.Printf("%d\t%b\t%#x",b,b,b)
}
