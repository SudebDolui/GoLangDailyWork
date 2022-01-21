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
	fmt.Println("Entered string:", my_string)

	for _, test_character := range my_string {

		if unicode.IsNumber(test_character) {
			check_number = true
		} else if unicode.IsLetter(test_character) {
			check_alphabet = true
		} else if unicode.IsSymbol(test_character) || unicode.IsSpace(test_character) {
			check_special_character = true
			fmt.Println(string(test_character))
		}
	}

	if check_alphabet && check_number && check_special_character == true {
		fmt.Println("The string contains Special, Numbers and Alphabets ")
	} else if (check_alphabet && check_number == true) && (check_special_character == false) {
		fmt.Println("The string does not contains special character")
	} else if (check_alphabet == true) && (check_special_character == false && check_number == false) {
		fmt.Println("The string does not contains special character and number")
	} else if (check_number == true) && (check_special_character == false && check_alphabet == false) {
		fmt.Println("The string does not contains special character and alphabets")
	} else {
		fmt.Println("The string does not contains specified characters")
	}
}
