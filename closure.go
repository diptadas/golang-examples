package main

import "fmt"

func main(){

	num := 2

	cls := func() int {
		num *= 2
		return  num
	}

	fmt.Println(num)
	fmt.Println(cls())
	fmt.Println(num)
	fmt.Println(cls())
	fmt.Println(num)

        fmt.Println("=============")

	cls2 := func(num int) int {
		num *= 2
		return  num
	}

	fmt.Println(num)
	fmt.Println(cls2(num))
	fmt.Println(num)
	fmt.Println(cls2(num))
	fmt.Println(num)
}
