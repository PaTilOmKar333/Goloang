package main

import "fmt"

func main() {

	omkar := User{"omkar", "omkar.patil@gmail.com", 27, true}
	fmt.Println(omkar)

	omkar.getstatus()

	omkar.setmail()

	fmt.Println(omkar)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getstatus() {
	fmt.Println(u.Status)
}

func (u User) setmail() {
	u.Email = "new@mail"
	fmt.Println(u.Email)
}
