package main

import (
	"fmt"
)

func lenghofstr(s string) int {
	couunter := 0
	for i, _ := range s {
		if string(s[i]) == " " {

		} else {
			couunter++
			//fmt.Println(s[i])
		}
	}
	return couunter
}

func main() {
	str := "i am omkar patil"
	fmt.Println(lenghofstr(str))
}
