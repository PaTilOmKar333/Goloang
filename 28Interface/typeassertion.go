package main

import "fmt"

func main() {
	var x interface{} = "omkar"

	//var y string = x.(string)

	C, OK := x.(string)
	fmt.Println(C, OK)

	D, OK := x.(int)
	fmt.Println(D, OK)

}
