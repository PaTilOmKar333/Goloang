package main

import (
	"fmt"
)

func main() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))
}

func itemExists(slicee []string, s string) bool {

	for i := 0; i < len(slicee); i++ {
		if slicee[i] == s {
			return true
		}
	}

	return false
}

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geomatric interface {
// 	area() float64
// 	perim() float64
// }

// type rect struct {
// 	height, width float64
// }

// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.height * r.width
// }

// func (r rect) perim() float64 {
// 	return 2*r.height + 2*r.width
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measurement(g geomatric) {
// 	fmt.Println(g)
// 	fmt.Println("area :", g.area())
// 	fmt.Println("perim :", g.perim())
// }

// func main() {

// 	r := rect{height: 4, width: 3}
// 	c := circle{radius: 5}

// 	measurement(r)
// 	measurement(c)
// }
