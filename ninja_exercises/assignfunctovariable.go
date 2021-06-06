package main

import "fmt"

func main(){
	foovalue:=assignfoo()
	foocheck:=func(){
		for i:=1;i<5;i++ {
			fmt.Println(i)
		}
	}
	foocheck()
	fmt.Println(foovalue)
	// Declare a varible and assign it value of func type
	variablefunctype:=returnfunc()
	// Call variable of func type
	variablefunctype()
	fmt.Printf("%T\n",variablefunctype)
}

func assignfoo() string{
	s:= fmt.Sprint("Hello this is foo calling...")
	return s
}

func returnfunc() func(){
	return func(){
		fmt.Println("hello")
	}
}