package main

import (
	"fmt"
	"time"
)

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}

func main() {
	f("omkar")
	go f("InGO1")
	go f("InGO2")

	go func(msg string) {
		fmt.Println(msg)
	}("adiraj")

	time.Sleep(time.Second)

	fmt.Println("Done")

}
