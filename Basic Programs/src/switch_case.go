package main

import "fmt"

// func main() {
// 	fmt.Println(1)
// 	fmt.Println(2)
// 	fmt.Println(3)
// 	fmt.Println(4)
// 	fmt.Println(5)
// 	fmt.Println(6)
// 	fmt.Println(7)
// 	fmt.Println(8)
// }

// switch expression {
// 	case value_1:
// 		statements_1
// 	case value_2:
// 		statements_2
// 	case value_n:
// 		statements_n
// 	default:
// 		statements_default
// }

func main() {
	var i int
	fmt.Println("Please Enter a Number : ")
	fmt.Scanf("%d", &i)
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}

// func main() {
// 	a, b := 2, 1
// 	switch a + b {
// 	case 1:
// 		fmt.Println("Sum is 1")
// 	case 2:
// 		fmt.Println("Sum is 2")
// 	case 3:
// 		fmt.Println("Sum is 3")
// 	default:
// 		fmt.Println("Printing default")
// 	}
// }
