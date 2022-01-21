package main

import (
	"fmt"
	"time"
)

//! Program 1:

// func displayString(ch chan string) {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Inside display()")
// 	ch <- "Hello"
// }

// func display(ch chan string, ch_one chan string) {
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Inside display()")
// 	ch <- "1234.123"
// 	ch_one <- "Hello123"
// }

// func main() {
// 	ch := make(chan string)
// 	ch_one := make(chan string)
// 	go display(ch, ch_one)
// 	go displayString(ch_one)
// 	x := <-ch
// 	y := <-ch_one
// 	fmt.Println("Inside main()")
// 	fmt.Println("Printing x in main() after taking from channel:", x, y)
// }

//! Program 2:

// You can also use channels for communication between goroutines.
// Need to use 2 goroutines – one pushes data to the channel and other receives the data from the channel.
// See the below program
//This subroutine pushes numbers 0 to 9 to the channel and closes the channel

// func add_to_channel(ch chan int) {
// 	fmt.Println("Send data")
// 	for i := 0; i < 10; i++ {
// 		ch <- i //pushing data to channel
// 	}
// 	close(ch) //closing the channel
// }

//This subroutine fetches data from the channel and prints it.

// func fetch_from_channel(ch chan int) {
// 	fmt.Println("Read data")
// 	for {
// 		//fetch data from channel
// 		x, flag := <-ch
// 		//flag is true if data is received from the channel
// 		//flag is false when the channel is closed
// 		if flag == true {
// 			fmt.Println(x)
// 		} else {
// 			fmt.Println("Empty channel")
// 			break
// 		}
// 	}
// }

// func main() {
// 	//creating a channel variable to transport integer values
// 	ch := make(chan int)
// 	//invoking the subroutines to add and fetch from the channel
// 	//These routines execute simultaneously
// 	go add_to_channel(ch)
// 	go fetch_from_channel(ch)
// 	//delay is to prevent the exiting of main() before goroutines finish
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Inside main()")
// }

//! Program 3:

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
