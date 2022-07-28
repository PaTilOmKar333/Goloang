package main

import (
	"fmt"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

//main func without waitgroups
// func main() {

// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	fmt.Println("Hello World")

// }

//main func with waitgroup
var x = 0

func main() {

	for i := 0; i < 10; i++ {
		go func(i int) {
			x = x + 1
			wg.Done()
		}(i)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(x)

}
