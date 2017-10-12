package main

import "fmt"

func main()  {
	var a [3]string
	a[0]="Helllo"
	a[1]="World"

	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes:=[7] int{2,3,4,5,7,11,13}
	fmt.Println(primes)
}
