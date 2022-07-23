package main

import (
	"fmt"
	"sort"
)

func main(){

	s := make([]int, 3)

	s[0]=134
	s[1]=562
	s[2]=3345

	s=append(s,743,856,9234)

	fmt.Println("s:",s)

	fmt.Println(sort.IntsAreSorted(s))
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))


	fmt.Println("s:",s)





}