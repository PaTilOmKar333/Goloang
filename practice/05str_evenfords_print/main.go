package main

import (
	"fmt"
	"strings"
)

func evenwords(str string) {
	a := strings.Split(str, " ")
	for _, v := range a {
		//fmt.Println(i, v)
		if len(v)%2 == 0 {
			fmt.Println(v)
		}
	}
}

func main() {
	str := "i will do it"
	evenwords(str)
}
