package main

import "fmt"

type User struct {
	Id int
	Age int
}

func call(obj interface{})  {
	fmt.Printf("%T %v\n", obj, obj)
}

func main () {
	call(User{1, 25})
}
