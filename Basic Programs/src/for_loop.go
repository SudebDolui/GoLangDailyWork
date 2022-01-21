package main

import "fmt"

// Program 1:

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("Welcome %d times.\n", i)
// 		// break out of for loop when i is 5
// 		if i == 5 {
// 			break
// 		}
// 	}
// }

// Program 2:

// func main() {
// 	var i int
// 	for i = 1; i <= 5; i++ {
// 		fmt.Println(i)
// 	}
// }

// Program 3:

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			fmt.Println(i, "even")
// 		} else {
// 			fmt.Println(i, "odd")
// 		}
// 	}
// }

// Program 4:

// func main() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i = i + 1
// 	}
// }

// Program 5:

func main() {
	// single condition for loop
	m := 1
	for m <= 5 {
		fmt.Printf("Welcome %d times.\n", m)
		m = m + 1
	}
	// classic for loop example
	for i := 6; i <= 10; i++ {
		fmt.Printf("Welcome %d times.\n", i)
	}
}
