package main

import "fmt"

func main() {
	const b = 10
	fmt.Println(b)
	// b =20 Error

	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	area = LENGTH * WIDTH
	fmt.Printf("value of area:%d", area)

}
