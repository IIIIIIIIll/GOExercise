package main

import "fmt"

func printSlice2(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main()  {
	var s  []int
	printSlice2(s)

	s=append(s,8)
	printSlice2(s)

	s=append(s,1)
	printSlice2(s)

	s=append(s,2,3,4)
	printSlice2(s)
}
