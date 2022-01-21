package main

// Program 1:

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	fmt.Println(total / float64(len(xs)))
// }

// This Program computes the average of a series of numbers.
// Finding the average like this is a very general problem.

// func average(xs []float64) float64 {
// 	panic("Not Implemented")
// }

// Program 2:

// func average(xs []float64) float64 {
// 	total := 0.0
// 	for _, v := range xs {
// 		total += v
// 	}
// 	return total / float64(len(xs))
// }

// Notice that we changed the fmt.Println to be a return instead.
// The return statement causes the function to immediately stop and return the value after it to the function that called this one.
// Modify main to look like this:

// func main() {
// 	xs := []float64{98, 93, 77, 82, 83}
// 	fmt.Println(average(xs))
// }

// Program 3:

// define function named total that accept int values and returns int

// func total(x int, y int) int {
// 	return x + y
// }

// func main() {
// 	// call our function and store result in the answer variable
// 	answer := total(10, 20)
// 	fmt.Println("10 + 20  = ", answer)
// 	fmt.Println("100 + 500 = ", total(100, 500))
// }

// Program 4:

//calc is the function name which accepts two integers num1 and num2
//(int, int) says that the function returns two values, both of integer type.

// func calc(num1 int, num2 int) (int, int) {
// 	sum := num1 + num2
// 	diff := num1 - num2
// 	return sum, diff
// }

// func main() {
// 	x, y := 15, 10
// 	//calls the function calc with x and y an d gets sum, diff as output
// 	sum, diff := calc(x, y)
// 	fmt.Println("Sum", sum)
// 	fmt.Println("Diff", diff)
// }

// Program 5:
