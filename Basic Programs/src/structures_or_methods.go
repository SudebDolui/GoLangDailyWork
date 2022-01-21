package main

import "fmt"

// Program 1:

//declared the structure named emp

// type emp struct {
// 	name    string
// 	address string
// 	age     int
// 	phone   uint32
// }

//Declaring a function with receiver of the type emp

// func (e emp) display() {
// 	fmt.Println("Name:", e.name)
// 	fmt.Println("Address:", e.address)
// 	fmt.Println("Age:", e.age)
// 	fmt.Println("Phone:", e.phone)
// 	fmt.Println(e.name, e.address, e.age)
// }

// func main() {
// 	//declaring a variable of type emp
// 	var empdata1 emp
// 	//Assign values to members
// 	empdata1.name = "John"
// 	empdata1.address = "Street-1, London"
// 	empdata1.age = 30
// 	//declaring a variable of type emp and assign values to members
// 	empdata2 := emp{
// 		"Raj", "Building-1, Paris", 25, 1234567891}
// 	//Invoking the method using the receiver of the type emp
// 	// syntax is variable.methodname()
// 	empdata1.display()
// 	empdata2.display()
// }

// Program 2:

// type Circle struct {
// 	x float64
// 	y float64
// 	r float64
// }

// Or

// type Circle struct {
// 	x, y, r float64
// }

// func main() {
// 	var c Circle
// 	// c := new(Circle)
// 	// c := Circle{x: 0, y: 0, r: 5}
// 	c = Circle{0, 0, 5}
// 	// We can access fields using the . operator:
// 	// fmt.Println(c.x, c.y, c.r)
// 	// c.x = 10
// 	// c.y = 5
// 	fmt.Println("Area of Circle:"circleArea(c))
// }

// Let's modify the circleArea function so that it uses a Circle:

// Struct Fields

// func circleArea(c Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// In main we have:
// c := Circle{0, 0, 5}
// fmt.Println(circleArea(c))

//change main to:
//   c := Circle{0, 0, 5}
//   fmt.Println(circleArea(&c))

// Struct Methods

// func circleArea(c Circle) float64 {
// 	return math.Pi * c.r * c.r
// }

// Program 3:

// type Rectangle struct {
// 	x1, y1, x2, y2 float64
// }

// func (r *Rectangle) area() float64 {
// 	l := distance(r.x1, r.y1, r.x1, r.y2)
// 	w := distance(r.x1, r.y1, r.x2, r.y2)
// 	return l * w
// }

// func main() {
// 	r := Rectangle{0, 0, 10, 10}
// 	fmt.Println(r)
// }

// Program 4:
// Embedded Types

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

func main() {
	a := new(Android)
	a.Person.Talk()
	// a := new(Android)
	// a.Talk

}
