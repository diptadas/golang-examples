package main

import "fmt"

func main(){

	fmt.Println(safeDivide(10, 5))
	fmt.Println(safeDivide(10, 0))
}

func safeDivide(num1, num2 int) int{

	defer func(){

		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	return  num1/num2
}
