package main

import "fmt"

func main() {
	var slice_1 = []int{3, 5, 76, 3, 6, 3, 5, 6, 3}
	var slice_2 = []int{2, 3, 65, 7, 4, 3, 6, 3, 56, 3}
	slice_3 := append(slice_1, slice_2...)
	fmt.Printf("slice_1: %v\n", slice_1)
	fmt.Printf("slice_2: %v\n", slice_2)
	fmt.Printf("Concantination on slice_1 and slice_2: %v\n", slice_3)
}
