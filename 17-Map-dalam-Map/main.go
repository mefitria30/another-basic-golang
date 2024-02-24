package main

import "fmt"

func main() {
	// student := map[string]string{
	// 	"1001": "ucil",
	// 	"1002": "timmy",
	// 	"1003": "molly",
	// }

	student := map[string]map[string]string{
		"1001": map[string]string{
			"name":    "ucil",
			"address": "indonesia",
			"gender":  "mail",
		},
		"1002": map[string]string{
			"name":    "timmy",
			"address": "indonesia",
			"gender":  "mail",
		},
		"1003": map[string]string{
			"name":    "molly",
			"address": "indonesia",
			"gender":  "mail",
		},
	}

	delete(student, "1003")

	fmt.Println(student)
	fmt.Println(student["1001"])
	fmt.Println(student["1001"]["name"])
	fmt.Println(student["1001"]["gender"])
}
