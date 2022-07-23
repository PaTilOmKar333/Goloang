package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	content := "hi i am omkar patil"

	file, err := os.Create("./test.txt")

	Errorfunc(err)

	length, err := file.WriteString(content)

	Errorfunc(err)

	fmt.Println(length)

	databyte, err := ioutil.ReadFile("./test.txt")

	Errorfunc(err)

	fmt.Println(databyte)
	fmt.Println(string(databyte))

}

func Errorfunc(err error) {

	if err != nil {
		panic(err)
	}
}
