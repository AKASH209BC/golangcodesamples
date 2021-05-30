package main

import "fmt"

// Create a program using showing if statement in action
func main(){
	my_byear:= 1996
	if my_byear > 2000 {
		fmt.Println(" I am a millenial")
	} else if my_byear > 1995 && my_byear <=2000 {
		fmt.Println("I was born in the late 90's")
	} else if my_byear > 1990 && my_byear <=1995 {
		fmt.Println("I was born in the early 90's")
	} else {
		fmt.Println("I am quite older now ")
	}
}
