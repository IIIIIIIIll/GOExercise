package main

import (
	"fmt"
)

type Vertex1 struct {
	x int
	y int
}

func main()  {
	v:=Vertex1{1,2}
	v.x=4
	fmt.Println(v.x)
}
