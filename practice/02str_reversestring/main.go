package main

import (
	"fmt"
	"strings"
)

func revers_string(str string) string {

	//a := []rune(str) // convert to rune
	a := strings.Split(str, " ")

	// for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	// 	a[i], a[j] = a[j], a[i]
	// }

	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	//fmt.Printf("%T", a)
	revers_str := strings.Join(a, " ")
	// fmt.Println(revers_str)
	// fmt.Printf("%T", revers_str)

	return revers_str
}

func main() {
	input_str := "omkar patil"
	reverseed_str := revers_string(input_str)
	fmt.Println(reverseed_str)
}
