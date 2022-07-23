package main

import "fmt"

func fibonecci(count int, c chan int) {
	x, y := 0, 1
	for i := 0; i < count; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	ch := make(chan int, 10)

	fibonecci(cap(ch), ch)

	for i := range ch {
		fmt.Println(i)
	}
}
