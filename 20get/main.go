package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev0909"

func main() {
	fmt.Println("lcp get request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	bytecode, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytecode))

}
