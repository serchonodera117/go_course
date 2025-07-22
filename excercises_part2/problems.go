package main

import (
	"fmt"
	"sync"
	"encoding/json"
	"net/http"
)


func main(){
	fmt.Println("--- Excercises ----------")

	// execute_excercise_1()
	// execute_excercise_2([]string{"hola", "hola", "perro", "oso", "panda", "panda", "hola"})
	// execute_excercise_3_multiple_goroutines(3)
	// execute_excercise_3_chanel()

	//excersise 4
	turn_server_on()
}

//Number one: declare a struct employee, fill up and iterate to receive its values
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


//Number two: map, decclarate a map [string]int that saves times that a word appears into a []String

func execute_excercise_2(arr []string){
	var mapWords = make(map[string]int)
	for _, value := range arr {
		
		if _,ok:= mapWords[value]; !ok {
			mapWords[value] = 1
		}else{
			mapWords[value] += 1
		}
	}

	fmt.Println("my map: ", mapWords)
}


//Number three: launch 5 coroutinas into a groupe
func execute_excercise_3_multiple_goroutines(n int){
	var groupToWait sync.WaitGroup

	for i:=0; i<n; i++{
		groupToWait.Add(1)
		go func (i int){
			defer groupToWait.Done()
			fmt.Println("Goroutine", i)
		}(i)
	}
	groupToWait.Wait()
	fmt.Println("Goroutines have finished")
}

func execute_excercise_3_chanel(){
	ch := make(chan string)
	go worker(ch)
	msg := <-ch       //#wait to receive something
	fmt.Println(msg)
}
func worker (ch chan string){
	ch <- "work ended"
}


func execute_excercise_4(w http.ResponseWriter, r *http.Request){
	response:= map[string]string{"message":"hello, world"}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func turn_server_on(){

	http.HandleFunc("/hello" , execute_excercise_4)
	http.ListenAndServe(":8080", nil)

}