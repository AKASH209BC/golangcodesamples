package main

import "fmt"

func main(){
	x:= []int{1,2,3,4,5,6}
	sumo := variadicfoo(x...)
	subpar := variadicbar(x)
	fmt.Println(sumo, subpar)
}

func variadicfoo(x ...int) int{
	total:=0
	for _,v:=range x{
		total+=v
	}
	return total
}

func variadicbar(x []int) int{
	total:=0
	for _,v:=range x{
		total+=v
	}
	return total
}
