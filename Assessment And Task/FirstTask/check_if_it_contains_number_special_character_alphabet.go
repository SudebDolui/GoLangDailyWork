package main

import (
	"fmt"
	"unicode"
)

func main() {
	var my_string string
	var check_number, check_alphabet, check_special_character bool

	fmt.Println("Enter a string:")
	fmt.Scanf("%s", &my_string)
	// my_string = "sudebd 1999.in@mouritech.com"
	fmt.Println("Entered string:", my_string)

	for _, test_character := range my_string {

		if unicode.IsNumber(test_character) {
			check_number = true
			// fmt.Println("Number", string(test_character))
		} else if unicode.IsLetter(test_character) {
			check_alphabet = true
			// fmt.Println("Alphabet", string(test_character))
		} else if test_character >= ' ' && test_character <= ')' || test_character >= '*' && test_character <= '/' || test_character >= ':' && test_character <= '@' || test_character >= '[' && test_character <= '`' || test_character >= '{' && test_character <= '~' {
			check_special_character = true
			// fmt.Println("Special Character", string(test_character))
		}
	}

	if check_alphabet && check_number && check_special_character {
		fmt.Println("The string contains Special Character, Numbers and Alphabets ")
	} else if (check_alphabet && check_number == true) && (check_special_character == false) {
		fmt.Println("The string does not contains Special Character")
	} else if (check_alphabet == true) && (check_special_character == false && check_number == false) {
		fmt.Println("The string does not contains Special Character and Number")
	} else if (check_number == true) && (check_special_character == false && check_alphabet == false) {
		fmt.Println("The string does not contains Special Character and Alphabets")
	} else {
		fmt.Println("The string does not contains specified conditions")
	}
}
