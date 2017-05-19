package main

import "fmt"

func merge (a []int, b []int) []int {

	var r = make([]int, len(a) + len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) { r[i+j] = a[i]; i++ }
	for j < len(b) { r[i+j] = b[j]; j++ }

	return r
}

func mergeSort (items []int, ch chan []int) {

	if len(items) < 2 {
		ch <- items
		return
	}

	ch1 := make(chan []int)
	ch2 := make(chan []int)

	var middle = len(items) / 2

	go mergeSort(items[:middle], ch1)
	go mergeSort(items[middle:], ch2)

	part1, part2 := <-ch1, <-ch2

	ch <- merge(part1, part2)
}

func main () {
	ch := make(chan []int)
	go mergeSort([]int{ 10, 9, 8, 4, 5, 6, 7, 3, 2, 1 }, ch)
	fmt.Println(<-ch)
}