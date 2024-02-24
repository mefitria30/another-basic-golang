package main

import "fmt"

func main() {
	names := [5]string{
		"timmy",
		"molly",
		"pussy",
		"ucil",
		"oren",
	}

	fmt.Println(names)

	slice1 := names[1:4]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[1:]
	fmt.Println(slice3)

	names[1] = "kucing"

	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice1[1] = "kucing_oren"

	fmt.Println(names)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
