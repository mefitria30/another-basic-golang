package main

import "fmt"

func main() {
	student := make(map[string]string)

	student["1001"] = "ucil"
	student["1002"] = "timmy"
	student["1003"] = "molly"

	fmt.Println(student)

	fmt.Println(student["1001"])
	fmt.Println(student["1002"])
	fmt.Println(student["1003"])

	for student_no, student_name := range student {
		fmt.Println("student_no:", student_no, "student_name:", student_name)
	}
}
