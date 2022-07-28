package main

import "fmt"

type student struct {
	name string
	id   int
}

func (s student) String() string {
	return fmt.Sprintf("name:%v and ID:%v", s.name, s.id)
}

func main() {
	s1 := student{name: "omkar", id: 1}
	s2 := student{name: "komal", id: 2}

	fmt.Println(s1)
	fmt.Println(s2)
}
