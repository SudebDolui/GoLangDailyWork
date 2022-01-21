package main

import "fmt"

// Go includes two built-in functions to assist with slices: append and copy.

// Program 1:

// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := append(slice1, 4, 5)
// 	fmt.Println(slice1, slice2)
// }

// After running this program slice1 has [1,2,3] and slice2 has [1,2,3,4,5].
// Append creates a new slice by taking an existing slice (the first argument) and appending all the following arguments to it.

// program 2:

// func main() {
// 	slice1 := []int{1, 2, 3}
// 	slice2 := make([]int, 2)
// 	copy(slice2, slice1)
// 	fmt.Println(slice1, slice2)
// }

// After running this program slice1 has [1,2,3] and slice2 has [1,2].
// The contents of slice1 are copied into slice2, but since slice2 has room for only two elements only the first two elements of slice1 are copied.

// Program 3:

// func main() {
// 	// declaring array
// 	a := [5]string{"one", "two", "three", "four", "five"}
// 	fmt.Println("Array after creation:", a)
// 	var b []string = a[1:4] //created a slice named b
// 	fmt.Println("Slice after creation:", b)
// 	b[0] = "changed" // changed the slice data
// 	fmt.Println("Slice after modifying:", b)
// 	fmt.Println("Array after slice modification:", a)
// }

// Program 4:

func main() {
	a := [5]string{"1", "2", "3", "4", "5"}
	slice_a := a[1:3]
	b := [5]string{"one", "two", "three", "four", "five"}
	slice_b := b[1:3]
	fmt.Println("Slice_a:", slice_a)
	fmt.Println("Slice_b:", slice_b)
	fmt.Println("Length of slice_a:", len(slice_a))
	fmt.Println("Length of slice_b:", len(slice_b))
	slice_a = append(slice_a, slice_b...) // appending slice
	fmt.Println("New Slice_a after appending slice_b :", slice_a)
	slice_a = append(slice_a, "text1") // appending value
	fmt.Println("New Slice_a after appending text1 :", slice_a)
	fmt.Println(a, b)
}
