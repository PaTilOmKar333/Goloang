package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("value of dice number is:", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("value is 1")
	case 2:
		fmt.Println("value is 2")
		fallthrough
	case 3:
		fmt.Println("value is 3")
		fallthrough
	case 4:
		fmt.Println("value is 4")
	case 5:
		fmt.Println("value is 5")
	case 6:
		fmt.Println("value is 6")
	}

}
