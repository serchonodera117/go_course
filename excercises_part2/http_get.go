package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Age   int     `json: age`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func (u User) LogIn() {
	fmt.Printf("%s has logged succesfully!", u.Email)
}

func main() {
	fmt.Println("GET REQUEST")
	get_user_list()
}

func get_user_list() {
	url := "http://127.0.0.1:8080/users"

	response, err := http.Get(url)
	if err != nil {
		panic(err)
		return
	}

	defer response.Body.Close()

	var userList = make(map[int]User)

	if err := json.NewDecoder(response.Body).Decode(&userList); err != nil {
		panic(err)
	}

	var chose int
	fmt.Println("users")
	for i, user := range userList {
		fmt.Println(i, ".-", "name: ", user.Name)
	}

	fmt.Println("Select your user: ")
	fmt.Scan(&chose)

	if _, ok := userList[chose]; !ok {
		fmt.Println("user", chose, " not found")
		return
	}

	userList[chose].LogIn()

	fmt.Println(" data: ", userList[chose])
	
	fmt.Println("original users: ", userList)
	
	fmt.Println(" sorted users: ", order_list(userList))
}

func order_list(myUMap map[int]User)[]User{
	var list []User
	
	for _,v := range myUMap {
		list = append(list, v)
	}

	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Age  <  list[j].Age {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}

	return list
}
