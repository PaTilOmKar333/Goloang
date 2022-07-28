package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("type :%T, value:%v\n", i, i)
}

func main() {
	s := "i am omkar patil"
	describe(s)

	i := 10
	describe(i)

	f := 34.34
	describe(f)

	a := struct {
		name string
	}{
		name: "omkar",
	}

	describe(a)
}
