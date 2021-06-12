package main

import (
	"fmt"
	"sort"
)

type user1 struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user1
type bylast []user1

func (a byAge) Len() int {
	return len(a)
}

func (a byAge)Swap(i,j int) {
	a[i],a[j]= a[j],a[i]
}

func (a byAge)Less(i,j int) bool{
	return a[i].Age<a[j].Age
}

func (l bylast) Len() int {
	return len(l)
}

func (l bylast)Swap(i,j int) {
	l[i],l[j]= l[j],l[i]
}

func (l bylast)Less(i,j int) bool{
	return l[i].Last<l[j].Last
}


func main() {
	u1 := user1{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user1{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user1{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user1{u1, u2, u3}

	sort.Sort(byAge(users))
	for _, v :=range users{
		fmt.Println(v.First,"\t",v.Last,"\t",v.Age)
		sort.Strings(v.Sayings)
		for _,x :=range v.Sayings{
			fmt.Printf("%v\n",x)
		}
		fmt.Println()
	}
	sort.Sort(bylast(users))
	for _, v :=range users{
		fmt.Println(v.First,"\t",v.Last,"\t",v.Age)
		sort.Strings(v.Sayings)
		for _,x :=range v.Sayings{
			fmt.Printf("%v\n",x)
		}
		fmt.Println()
	}
}
