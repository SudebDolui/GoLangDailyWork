package main

import "fmt"

// Program 1:

// func main() {
// 	var x map[string]int
// 	x["key"] = 12
// 	fmt.Println(x)
// 	delete(x, "key")
// 	fmt.Println(x)
// }

// Program 2:

func main() {
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"])
	// If you run this you should see nothing returned.
	// Technically a map returns the zero value for the value type (which for strings is the empty string).
	// Although we could check for the zero value in a condition (elements["Un"] == "")
	name, ok := elements["Un"]
	fmt.Println(name, ok)
	// Accessing an element of a map can return two values instead of just one.
	// The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
