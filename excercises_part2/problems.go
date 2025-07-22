package main

import "fmt"


func main(){
	fmt.Println("--- Excercises ----------")

	execute_excercise_1()
}

//Number one
func execute_excercise_1(){
	var employee_slice = []Employee{{ID: 1, Name: "Juan", Position: "Recruiter"}}

	employee_slice = append(employee_slice ,  Employee{ID: 2, Name: "Lorenso", Position: "Dev"})
	employee_slice = append(employee_slice , Employee{ID: 3, Name: "Leopold", Position: "Assitant"})

	for index, value := range employee_slice {
		fmt.Println("i:", index, " content:", value)
	}
}

type Employee struct {
	ID int
	Name string
	Position string
}