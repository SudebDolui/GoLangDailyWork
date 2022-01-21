package main

import (
	"com.table/packages/tablegenerator"
	"fmt"
)

func main() {
	var number, upto int
	fmt.Print("Enter the number for which you want the table :")
	fmt.Scanf("%d\n", &number)
	fmt.Print("Upto which number you want the table :")
	fmt.Scanf("%d\n", &upto)
	tablegenerator.TableGenerator(number, upto)

}
