package main

import "fmt"

func main(){
	a:= value1()
	b:= value1()
	c:= value1()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(c())

}


func value1() func() int{
	var y int
	return func() int {
		for i := 0; i < 2; i++ {
			y++
		}
		return y
	}
}