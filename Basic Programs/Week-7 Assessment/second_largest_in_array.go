package main

import "fmt"

func main() {
	var large1 int = 0
	var large2 int = 0
	arr := [...]int{3, 5, 76, 3, 6, 3, 5, 6, 3}

	large1 = arr[0]
	for i := 1; i <= 4; i++ {
		if large1 < arr[i] {
			large2 = large1
			large1 = arr[i]
		} else if large2 < arr[i] {
			large2 = arr[i]
		}
	}
	fmt.Println("Second largest element is: ", large2)
}
