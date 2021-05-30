package main

import "fmt"

// Create an array of size 5 , assign values to all indexes , print them and the type of the array
func main(){
	// Empty int array ; array becomes 0, 0, 0, 0, 0 as 0 is the zero (default) value for int type
	x:= [5]int{}
	// Print out the entire array
	fmt.Println(x)
	x[0]=1
	x[1]=2
	x[2]=3
	x[3]=4
	x[4]=5
	// Print specific indexes
	fmt.Println("The value at index 1 is ", x[1])
	fmt.Println("The value at index 2 is ", x[2])
	// Print value using range function
	for i, v := range x {
		fmt.Println(i,v)
	}
	fmt.Printf("%T\n",x)
}
