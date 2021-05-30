package main

import "fmt"

func main(){

	// Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map.
	// Print out all of the values, along with their index position in the slice.
	//	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`
	x:= map[string][]string{}
	x["bond_james"] = []string{"Shaken, not stirred","Martinis","Women"}
	x["moneypenny_miss"] =[]string{"James Bond","Literature","Computer Science"}
	x["no_dr"]=[]string{"Being evil","Ice cream","Sunsets"}
	for i, v := range x {
		fmt.Printf("The key is %v\n",i)
		for j,k := range v {
			fmt.Printf("The value at index %d is %v\n", j, k)
		}
	}
	// Add a record to the map
	// Print out the map using range keyword
	x["bharat"]=[]string{"All is Well","ice cream","martini"}
	for i,v :=range x {
		fmt.Printf("The key is %v\n",i)
		for j,k := range v {
			fmt.Printf("The value at index %d is %v\n", j, k)
		}
	}
	// Delete a value from the map
	// Print using range loop
	if seconds, ok := x["bond_james"]; ok {
		fmt.Printf("Deleting the value %v\n",seconds)
		delete(x,"bond_james")
		fmt.Println(x)
	}
	//for i,ok:=x["bond_james"];ok{
	//}
}
