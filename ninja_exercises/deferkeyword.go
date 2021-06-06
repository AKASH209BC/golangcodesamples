package main

import "fmt"

func main(){
	fmt.Println("Starting defoo execution... ")
	defer defoo()
	fmt.Println("Defer encountered.. Will wait for enclosing function i.e. main to complete before executing defoo again..")
	debar()
	fmt.Println("Executing debar function")
	fmt.Println("End of main encountered... Starting defoo execution again")
}

func defoo(){
	fmt.Println("Executing foo function")
}

func debar(){
	fmt.Println("Executing bar function")
}
