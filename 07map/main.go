package main

import "fmt"

func main(){

	langauge := make(map[int]string)
	langauge[1]="omkar"
	langauge[2]="komal"
	langauge[3]="patil"

	fmt.Println("langauge:",langauge)


	for _, value := range langauge{
		fmt.Printf("value %v \n",value)

	}


}
