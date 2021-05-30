package main

import "fmt"

func main(){
	x:= 33
	for x<140{
		fmt.Printf("%d\t%#U\n",x,x)
		x++
	}
}
