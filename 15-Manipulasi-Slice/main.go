package main

import "fmt"

func main() {
	slice1 := make([]string, 3)
	slice1[0] = "molly"
	slice1[1] = "pussy"
	slice1[2] = "ucil"

	slice2 := append(slice1, "bunga", "mawar")

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice1[0] = "melati"

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]string, 3)

	copy(slice3, slice1)

	fmt.Println(slice1)
	fmt.Println(slice3)

	slice1[0] = "rose"

	fmt.Println(slice1)
	fmt.Println(slice3)
}
