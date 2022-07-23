package main

import "fmt"

func sum(sa []int, ch chan int) {
	sum := 0

	for _, v := range sa {
		sum += v
	}
	ch <- sum
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
