package main

import (
	"fmt"
	"strings"
)

func capitalHalfstring(str string) {
	l := len(str) - 1
	//mid := l / 2
	var rstr string
	for i, _ := range str {
		if i == 0 && i == l {
			rstr = rstr + strings.ToUpper(string(str[i]))
			//	rstr = rstr + Strings(string(str[i])).ToUpper
		} else {
			rstr = rstr + string(str[i])
		}
	}
	fmt.Println(rstr)

}
func main() {
	str := "omkarpatil"
	capitalHalfstring(str)
}
