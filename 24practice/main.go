package main

import (
	"fmt"
)

func main() {
	var Name string
	fmt.Println("Enter string:")
	fmt.Scanln(&Name)

	//var newNAME string

	for _, c := range Name {
		//fmt.Println(i, c)
		//	fmt.Printf("%d %c\n", i, c)
		if c > 96 {
			//fmt.Printf("%c", c)
			c -= 32
			fmt.Printf("%c", c)
		} else if c >= 65 && c < 97 {
			//fmt.Printf("%c", c)
			c += 32
			fmt.Printf("%c", c)
		}

	}

}
