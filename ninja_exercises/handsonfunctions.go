package main

import "fmt"

func main(){
	x := foo()
	y,z := bar()
	fmt.Println("The value of x returned from func foo is", x)
	fmt.Println("The value of y and z returned from func bar are",y, "and", z, "respectively")

}

func foo() int{
	return 1996
}

func bar() (string,bool){
	s:= fmt.Sprint("Hello this is func bar calling")
	return s,true
}
