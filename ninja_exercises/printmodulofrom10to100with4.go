package main
import "fmt"

// Print remainder of each number from 10 to 100 when dividing by 4

func main(){

	for x:=10;x<=100;x++{
		fmt.Printf("%d\n",x%4)
	}

}
