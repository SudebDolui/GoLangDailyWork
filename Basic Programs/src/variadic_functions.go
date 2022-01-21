package main

import "fmt"

// There is a special form available for the last parameter in a go function

func main() {
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
