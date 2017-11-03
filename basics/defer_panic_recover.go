// https://blog.golang.org/defer-panic-and-recover

package main

import "fmt"

func main() {
	fmt.Println("main: Call c")
	c()
	fmt.Println("main: Return c")

	fmt.Println("main: Call f")
	f()
	fmt.Println("main: Return f")
}

func c() {
	defer fmt.Println("c: Defer 1")
	defer fmt.Println("c: Defer 2")

	fmt.Println("c: Return")
	return

	defer fmt.Println("c: Defer 3")
}

func f() {
	defer func() {
		defer fmt.Println("f: Defer 1")
		if r := recover(); r != nil {
			fmt.Println("f: Recover 1", r)
		}
	}()
	defer func() {
		defer fmt.Println("f: Defer 2")
		if r := recover(); r != nil {
			fmt.Println("f: Recover 2", r)
		}
	}()

	fmt.Println("f: Call g")
	g()
	fmt.Println("f: Return g")
}

func g() {
	defer fmt.Println("g: Defer 1")
	defer fmt.Println("g: Defer 2")

	fmt.Println("g: Panic")
	panic(fmt.Sprintf("g: Panic"))

	defer fmt.Println("g: Defer 3")
}
