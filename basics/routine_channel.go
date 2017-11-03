package main

import "fmt"
import "time"

func f1(str chan string) {
	fmt.Println("working 1")
	time.Sleep(time.Second * 3)
	str <- "1"
}

func f2(str chan string) {
	fmt.Println("working 2")
	time.Sleep(time.Second * 2)
	str <- "2"
}

func main() {

	str := make(chan string, 2)

	go f1(str)
	go f2(str)

	for i := 0; i < 2; i++ {

		fmt.Println("waiting for msg")
		msg := <-str

		if msg == "1" {
			fmt.Println("done 1")
		} else if msg == "2" {
			fmt.Println("done 2")
		}
	}

	fmt.Println("done 1 2")
}
