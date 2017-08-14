package main

//resource: https://golang.org/pkg/fmt/

import "fmt"

type point struct {
	x, y int
}

func main() {

	fmt.Println("Hello World")
	fmt.Println("sum of", 2, "and", 3, "is", 2+3)

	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p) //includes the field names
	fmt.Printf("%T\n", p)  //prints the type

	fmt.Printf("%b\n", 15) //binary
	fmt.Printf("%x\n", 15) //hex

}
