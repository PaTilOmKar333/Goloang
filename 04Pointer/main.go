package main

import "fmt"

func main(){
	var ptr *int

	myNumber:=23

	ptr = &myNumber

	fmt.Println("ptr:",ptr)
	fmt.Println("*ptr",*ptr)
}