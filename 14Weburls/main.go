package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	bytecode, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(bytecode))
}
