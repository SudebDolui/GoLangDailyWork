package main

import (
	"fmt"
	"time"
)

//! Program 1:

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 	}
// }

// func main() {
// 	go f(0)
// 	var input string
// 	fmt.Scanln(&input)
// }

//! Program 2:

// func display() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("In display")
// 	}
// }

// func main() {
// 	//invoking the goroutine display()
// 	go display()
// 	//The main() continues without waiting for display()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("In main")
// 	}
// }

//! Program 3:

// func f(n int) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(n, ":", i)
// 		amt := time.Duration(rand.Intn(250))
// 		time.Sleep(time.Millisecond * amt)
// 	}
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go f(i)
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }

//! Program 4:

func display() {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("In display")
	}
}

func main() {
	//invoking the goroutine display()
	go display()
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println("In main")
	}
}
