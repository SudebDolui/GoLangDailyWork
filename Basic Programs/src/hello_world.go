// package main

// import "fmt"

// func main() {

// 	fmt.Println("Welcome to Go Programming Language ...😀")

// }

//  OR

package main

// Import OS and fmt packages
import (
	"fmt"
	"os"
)

// import ("fmt";"os")

// Let us start
func main() {
	fmt.Println("Hello, world!")                           // Print simple text on screen
	fmt.Println(os.Getenv("USERS"), ", Let's be friends!") // Read Linux $USER environment variable
}
