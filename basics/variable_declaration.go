package main

import "fmt"

func main() {

	var str string = "hello"
	fmt.Println(str)

	var str2 = "hello" //infer data type based on value
	fmt.Println(str2)

	str3 := "hello"
	fmt.Println(str3)

	//declaring and initializing multiple variables

	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	p, q, r := 1, 2, 3
	fmt.Println(p, q, r)

	var (
		a = 5
		b = 10.2
		c = "string"
	)
	fmt.Println(a, b, c)

	//constants
	const pi float32 = 3.1416
	const pi2 = 9.87 //infer data type based on value
	fmt.Println(pi, pi2)

	const (
		_pi  = 3.1416
		_pi2 = 9.87
	)
	fmt.Println(_pi, _pi2)
}
