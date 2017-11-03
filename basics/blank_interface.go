package main

import "fmt"

type user struct {
	id  int
	age int
}

func call(obj interface{}) {
	fmt.Printf("%T %v\n", obj, obj)
}

func main() {
	call(365)
	call(3.14)
	call("hello")
	call(user{1, 25})
}
