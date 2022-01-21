package main

import "fmt"

func main() {
	var user_string string
	var check_for_vowel bool
	fmt.Print("Enter a String to check if vowel is present or not:")
	fmt.Scanf("%s\n", &user_string)
	fmt.Println("Entered string is ", user_string)

	for _, char := range user_string {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			check_for_vowel = true
		}
	}

	if check_for_vowel {
		fmt.Println("The String Contains vowel.")
	} else {
		fmt.Println("There are No vowel in the String.")
	}

}
