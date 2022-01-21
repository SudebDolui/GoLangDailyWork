package main

import (
	"fmt"
	"unicode"
)

func main() {
	var my_string string
	fmt.Println("Enter a String:")
	fmt.Scanf("%s\n", &my_string)

	for _, my_character := range my_string {

		if unicode.IsNumber(my_character) {
			// fmt.Printf("key[%d] = %d \n", i, my_character)
			fmt.Println(string(my_character), "With Unicode", my_character, "It's a Number.")
		} else if unicode.IsLetter(my_character) {
			fmt.Println(string(my_character), "With Unicode", my_character, "It's an alphabet.")
		} else {
			fmt.Println(string(my_character), "With Unicode", my_character, "It's neither an Alphabet or a number.")
		}

	}

}
