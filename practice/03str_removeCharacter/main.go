package main

import "fmt"

func removecharacter(i int, str string) string {
	var newstr string
	for j, _ := range str {
		if j != i {
			//fmt.Println("j ", j, string(str[i]), v, str[i])
			//fmt.Println(string(str[i]))
			newstr = newstr + string(str[j])
		}

	}
	//fmt.Println(newstr)

	return newstr
}

func main() {
	place := 3
	str := "omkar"
	newstr := removecharacter(place, str)
	fmt.Println(newstr)
}
