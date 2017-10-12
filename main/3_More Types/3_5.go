package main

import "fmt"

type Vertex3 struct {
	x,y int
}

var (
	v1 = Vertex3{1,2}
	v2=Vertex3{x:1}
	v3=Vertex3{}
)


func main()  {
	p:=&v1
	fmt.Println(v1,*p,v2,v3)
}