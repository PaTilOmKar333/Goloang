package main

import "fmt"

func main() {
	v1, _ := supperadder(2, 3, 4)
	fmt.Println(v1)
	v2, m2 := supperadder(2, 3)
	fmt.Println(v2)
	fmt.Println(m2)
	v3, _ := supperadder(2, 3, 4, 5)
	fmt.Println(v3)
}

func supperadder(values ...int) (int, string) {
	total := 0

	for _, value := range values {
		total += value
	}
	return total, "hi addition done"
}
