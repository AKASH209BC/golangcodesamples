package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main(){
	truck1:= truck{
		vehicle{
			doors: 4,
			color: "blue",
		},
		true,
	}
	sedan1:= sedan{
		vehicle{
			doors: 4,
			color: "white",
		},
		true,
	}
	fmt.Printf("The value of truck is %v\n",truck1)
	fmt.Printf("The value of sedan is %v\n",sedan1)
	fmt.Println(truck1.fourWheels,truck1.doors,truck1.color)
	fmt.Println(sedan1.luxury,sedan1.doors,sedan1.color)

}
