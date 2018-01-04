package main

import (
	"fmt"
	"log"

	"github.com/google/go-querystring/query"
)

// Encoding structs into URL query parameters
// Ref: https://github.com/google/go-querystring

type Person struct {
	Name    string  `url:"name"`
	Age     int     `url:"age"`
	Address Address `url:"address"`
}

type Address struct {
	City     string `url:"city"`
	PostCode int    `url:"postcode"`
}

func main() {
	opt := Person{
		Name: "Dipta",
		Age:  10,
		Address: Address{
			City:     "Chittagong",
			PostCode: 4000,
		},
	}

	value, err := query.Values(opt)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(value)
	fmt.Print(value.Encode())
}
