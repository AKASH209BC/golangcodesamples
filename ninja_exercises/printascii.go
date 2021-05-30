package main

import "fmt"

func main(){
	x:= 33
	for x<140{
		fmt.Printf("%d\t%#U\n",x,x)
		x++
	}
	// Print below Pattern for capital alphabets using ASCII values
	// 65 - A's ASCII Code
	// U+0041 'A' - 3 times in 3 lines
	// new line
	// Repeat above for all capital alphabets

	for i:=65;i<91;i++{
		fmt.Println(i)
		for j:=0;j<3;j++ {
			fmt.Printf("%#U\t\n",i);
		}
		fmt.Printf("\n")
	}

	// Print Years you have been alive using for condition {} syntax
	year:=1996
	for year<2022{
		fmt.Println(year)
		year++
	}

	// Print Years you have been alive using for for {} syntax
	year:= 1996
	for {
		if year>2021{
			break
		}
		fmt.Println(year)
		year++
	}

}
