package main

import (
	"fmt"
)

func palandrom(str string) {
	mid := len(str) / 2
	start := 0
	last := len(str) - 1
	flag := 0
	fmt.Println(mid, start, last)
	for start < last {
		if str[start] == str[last] {
			start += 1
			last -= 1
		} else {
			flag = 1
			break
		}
	}
	if flag == 1 {
		fmt.Println("this is not palandrom:", flag)
	} else {
		fmt.Println("this is palandrom:", flag)
	}

}

func symmetry(str string) {
	var mid int
	length := len(str)
	flag := 0
	fmt.Println("length", length)
	if length%2 == 0 {
		mid = length / 2
		fmt.Println("mid", mid)
	} else {
		mid = length/2 + 1
		fmt.Println("mid", mid)
	}
	start1 := 0
	start2 := mid
	for start1 < mid && start2 < length {
		fmt.Println("str1", str[start1], start1)
		fmt.Println("str2", str[start2], start2)
		if str[start1] == str[start2] {
			fmt.Println("str1", str[start1])
			fmt.Println("str2", str[start2])
			start1 += 1
			start2 += 1
		} else {
			flag = 1
			break
		}
	}
	if flag == 0 {
		fmt.Println("this is symmetry:", flag)
	} else {
		fmt.Println("this is not symmetry:", flag)
	}

}

func main() {

	str := "akaaka"
	//palandrom(str)
	symmetry(str)
}
