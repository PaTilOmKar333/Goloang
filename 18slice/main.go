package main

import "fmt"

func main() {

	slice1 := make([]int, 5)

	fmt.Println(slice1, len(slice1), cap(slice1))

	slice2 := make([]int, 0, 5)

	fmt.Println(slice2, len(slice2), cap(slice2))

	slice3 := slice2[:2]

	fmt.Println(slice3, len(slice3), cap(slice3))

	slice4 := slice3[2:5]
	fmt.Println(slice4, len(slice4), cap(slice4))

}
