package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("---- Writing and Reading Files -----")
	//writing persone
	// var person  Persona = Persona{
	// 	Name: "Serch",
	// 	Age: 24,
	// 	Email: "cherchin@email.com ",
	// }

	// write_file(person)

	read_file()
}

type Persona struct {
	Name string   `json: "name"`
	Age int    `json: "age"`
	Email string  `json: "email"`
}

func write_file(data Persona){
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err!= nil{
		panic(err)
	}

	err = os.WriteFile("persona.json", jsonData, 0644)
	if err != nil{
		panic(err)
	}

	fmt.Println("JSON filw ritten succesfully!")
}

func read_file(){
	file, err := os.Open("persona.json")
	if err!=nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err!= nil{
		panic(err)
	}

	var tempPerson Persona
	err = json.Unmarshal(byteValue, &tempPerson)

	fmt.Printf("Name: %s, Age: %d, Email: %s", tempPerson.Name, tempPerson.Age, tempPerson.Email)
}
