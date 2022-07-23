package main

import "fmt"

func main(){

	var fruitList [4]string

	fruitList[0]="Apple"
	fruitList[1]="pinaple"
	// fruitList[2]="crustedapple"
	fruitList[3]="maongo"

	fmt.Printf("type of array:%T\n",fruitList)

	fruitList=append(fruitList,"cadburry")


	fmt.Println("fruits:",fruitList)

	
	veggies:=[5]string{"potatao1","potatao2","potatao3","potatao4"}

	fmt.Println("veggies",veggies)
}