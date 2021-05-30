package main
import "fmt"
//With the help of fallthrough statement, we can use to transfer the program control just after the statement is executed in the switch cases even if the expression does not match.
//Normally, control will come out of the statement of switch just after the execution of first line after match.
//Don’t put the fallthrough in the last statement of switch case.
func main(){
	x:=12
	// switch without condition
	switch{
	case x>12: fmt.Printf("%d is not greater than 12\n",x)
	case x==12: fmt.Printf("%d is equal to 12\n",x)
	case x<12: fmt.Printf("%d is lesser than 12\n",x)
	default: fmt.Printf("Default statement\n")
	}
	// switch with condition
	favSport:= "Cricket"
	switch favSport{
	case "Cricket": fmt.Println("I love Cricket")
	case "Football": fmt.Println("I love Football")
	default:
		fmt.Println("I admire all sports and don't love any")
	}
}
