package main

import "fmt"

func main() {
	var hello string = "Hello World"

	fmt.Println(hello)

	hello = "Hello Golang"

	fmt.Println(hello)

	fullName := "Fitria Melliyanni" // by default, varchar
	fmt.Println(fullName)

	grade := 10 // by default, integer
	fmt.Println(grade)

	const flower string = "rose"
	fmt.Println(flower)

	// flower = "jasmine" // cannot reassign on constant data type
}
