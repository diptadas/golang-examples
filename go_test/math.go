package test

func Sum(values []int) int {
	sum := 0
	for _, num := range values {
		sum += num
	}
	return sum
}
