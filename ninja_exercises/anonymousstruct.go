package main

import "fmt"

type sects struct{
	religions []string
	countries map[string]string
}
type value struct{
	sects
}
func main(){
	sect1 := sects{
		[]string{
			"Hinduism",
			"Islam",
		},
		map[string]string{
			"India":"New Delhi",
			"Pakistan":"Islamabad",
		},
	}
	sect2 := sects{
		[]string{
			"Christianity",
			"Sikhism",
		},
		map[string]string{
			"India":"UK",
			"Pakistan":"Canada",
		},
	}
	fmt.Println("The religions in the sects are ")
	for x:=0;x<2;x++{
		fmt.Printf("%v and %v\n",sect1.religions[x],sect2.religions[x])
	}
	fmt.Println("The countries in the sects are ")
	for i,v := range sect1.countries{
		fmt.Printf("%v and %v\n",i,v)
	}
}
