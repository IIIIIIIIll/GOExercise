package main

import (
	"fmt"
)

type Verrtex2 struct {
	x int
	y int
}

func main()  {
	v:=Verrtex2{1,2}
	p:=&v
	p.x=1e9
	fmt.Println(v)
}
