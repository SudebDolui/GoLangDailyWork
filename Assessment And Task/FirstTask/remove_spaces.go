package main

import "fmt"

func main() {
	var user_input string
	var corrected_output []string
	fmt.Println("Enter a word in which you want to remove spaces:")
	// fmt.Scanf("%s", &user_input)
	user_input = "S U M"

	for _, char := range user_input {
		i := 0
		if char != ' ' {
			corrected_output[i] = string(char)
			i++
		}
	}

	fmt.Println("Corrected Output after removing spaces is", corrected_output)
}
