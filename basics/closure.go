package main

import "fmt"

func main() {
	num := 0

	cls := func(n int) int {
		fmt.Println("closure cls called")
		if num > n {
			return num - n
		} else {
			return n - num
		}
	}

	func(n int) {
		fmt.Println("anonymous closure called")
		num = num + n
	}(10)

	fmt.Println(num)
	fmt.Println(cls(100))
	fmt.Println(cls(500))
}
