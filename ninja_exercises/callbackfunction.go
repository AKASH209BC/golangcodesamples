package main

import "fmt"

func main(){
	x:=[]int{2,4,9,11,23}
	multiplyByThree(printOddValues,x...)
}


func printOddValues(x ...int ){
	for _,v := range x{
		if v%2!=0 {
			fmt.Println(v)
		}
	}
}

func multiplyByThree(p func (x ...int), y ...int){
	var x [] int
	for _,v :=range y{
		v*=3
		x= append(x,v)
	}
	p(x...)
}