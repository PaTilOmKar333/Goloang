package main

import (
	"fmt"
)

func main() {

	datachan := make(chan int)

	datachan <- 999

	n := <-datachan

	fmt.Println(n)
}
