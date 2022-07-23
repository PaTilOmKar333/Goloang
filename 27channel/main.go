package main

import "fmt"

func main() {

	message := make(chan string)

	go func() {
		message <- "dhanraj"
	}()

	msg := <-message
	fmt.Println(msg)
}
