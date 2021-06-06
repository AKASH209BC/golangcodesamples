package main

import "fmt"

func main(){

	fmt.Println("Using functions without identifier names ")
	func(x int){
		fmt.Println("The value is", x)
	}(42)

}
