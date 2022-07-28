package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	omkar := User{"Omkar", "omkar.patil@gmail.com", true, 25}

	// komal := User{
	// 	Name:   "komal",
	// 	Email:  "komal.com",
	// 	Status: true,
	// 	Age:    22,
	// }

	fmt.Println("omkar :", omkar)
	fmt.Printf("Name:%v, Email:%v, Status:%v, Age:%v\n", omkar.Name, omkar.Email, omkar.Status, omkar.Age)
}
