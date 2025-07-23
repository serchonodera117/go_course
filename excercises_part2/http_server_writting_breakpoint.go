package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var resp = make(map[int]User)

func main() {
	fmt.Println("EXCERCISE WITH SERVER")

	turn_server_on()
}

type User struct {
	ID    int    `json:"id"`
	Age   int     `json: age`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func (u User) LogIn() {
	fmt.Printf("%s has logged succesfully!", u.Email)
}

func sendUserList(w http.ResponseWriter, r *http.Request) {
	response := map[int]User{
		1: User{ID: 1, Age: 21, Name: "Serch", Email: "Sercho@gmail.com"},
		2: User{ID: 2, Age: 22, Name: "Porongol", Email: "p_onglo@gmail.com"},
		3: User{ID: 3, Age: 18, Name: "Belle", Email: "bellecchi@gmail.com"},
		4: User{ID: 4, Age: 28, Name: "Nicole", Email: "nicole@gmail.com"},
		5: User{ID: 5, Age: 58, Name: "a", Email: "a@gmail.com"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func turn_server_on() {
	http.HandleFunc("/users", sendUserList)
	http.ListenAndServe(":8080", nil)
}
