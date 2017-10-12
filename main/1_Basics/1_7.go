package main

import (
	"fmt"
)

func split(a int) (x, y int) {
	x=a*4/9
	y=a-x
	return
}

func main()  {
	fmt.Println(split(17))
}