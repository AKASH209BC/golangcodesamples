package main

import "fmt"

// Create an slice with five elements 10 , assign values to all indexes , print them and the type of the array
func main(){
	x:=[]int{42,43,44,45,46,47,48,49,50,51}

	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i,v)
	}

	fmt.Printf("%T\n",x)

	// Use Slicing to create the following slices

	// [42,43,44,45,46]
	// [47,48,49,50,51]
	// [44,45,46,47,48]
	// [43,44,45,46,47]
	fmt.Println("The first slice is ", x[:5])
	fmt.Println("The second slice is ", x[5:])
	fmt.Println("The third slice is ", x[2:7])
	fmt.Println("The fourth slice is ", x[1:6])

	// Append value 52 to slice x
	// Print the slice
	// In one statement append 53,54,55 to slice
	// Print the slice
	// Append slice with elements 56,...., 60 to slice x
	// Print the slice
	x= append(x,52)
	fmt.Printf("After adding 52 to slice x the slice becomes %v and its length %d\n",x,len(x))
	x= append(x,53,54,55)
	fmt.Printf("After adding 53,54,55 to slice x the slice becomes %v and its length %d\n",x,len(x))
	y:=[]int{56,57,58,59,60}
	x= append(x,y...)
	fmt.Printf("After adding slice y with values %v and length %d to slice x the slice becomes %v and its length %d\n",y, len(y), x,len(x))

	// Start with slice z equal to 42,43,44, ... , 51
	// Delete elements from it to make it 42,43,44,48,49,50,51
	z:=[]int{42,43,44,45,46,47,48,49,50,51}
	z= append(z[:3], z[6:]...)
	fmt.Println(z)

	// Create a slice to store the names of all of the states in the United States of America.
	// What is the length of your slice? What is the capacity?
	// Print out all of the values, along with their index position in the slice, without using the range clause.
	statenames := make([]string, 50,50)
	statenames=[]string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(statenames)
	for i:=0;i<=49;i++{
		fmt.Printf("The value at index %d is %v\n",i,statenames[i])
	}

	//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	//"James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	// Range over the records, then range over the data in each record.
	fv:=[]string{"James","Bond","Shaken not stirred"}
	sv:=[]string{"Miss", "Moneypenny", "Helloooooo, James."}
	md:=[][]string{fv,sv}
	for i,v := range md {
		fmt.Printf("At index %d for slice of slices \n",i)
		for j,k := range v{
			fmt.Printf("the value at index %d inside the slice is %v\n",j,k)
		}
	}

}
