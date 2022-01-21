package tablegenerator

import "fmt"

func TableGenerator(number int, upto int) {
	var Result int
	for i := 0; i <= upto; i++ {
		Result = number * i
		fmt.Println(number, "*", i, "=", Result)
	}
}
