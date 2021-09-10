package ginkgo_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Flow", func() {
	BeforeEach(func() {
		fmt.Println("Before each - 0")
	})

	JustBeforeEach(func() {
		fmt.Println("Just Before each - 0")
	})

	AfterEach(func() {
		fmt.Println("After each - 0")
	})

	Describe("Describe - 0", func() {
		Context("Context - 00", func() {
			BeforeEach(func() {
				fmt.Println("Before each - 000")
			})

			JustBeforeEach(func() {
				fmt.Println("Just Before each - 000")
			})

			AfterEach(func() {
				fmt.Println("After each - 000")
			})

			It("It - 000", func() {
				fmt.Println("It - 000")
			})
		})

		Context("Context - 01", func() {
			It("It - 010", func() {
				fmt.Println("It - 010")
			})

			It("It - 011", func() {
				fmt.Println("It - 011")
			})
		})
	})

})
