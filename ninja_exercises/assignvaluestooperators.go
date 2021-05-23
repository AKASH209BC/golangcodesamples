package main

import "fmt"

var a bool = 32==12
var b bool = 32!=12
var c bool = 32>=12
var d bool = 32!=12
var e bool = 32>12
var f bool = 32<12

func main(){
	fmt.Println("value of a is ", a)
	fmt.Println("value of b is ", b)
	fmt.Println("value of c is ", c)
	fmt.Println("value of d is ", d)
	fmt.Println("value of e is ", e)
	fmt.Println("value of f is ", f)
}