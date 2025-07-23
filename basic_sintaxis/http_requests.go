package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"

	// "io"
	"net/http"
	// "io/ioutil"
)

func main(){
	fmt.Println("--- HTTP RQUESTS---")
		var code string
	
	for code != "q" && code != "Q" {
		fmt.Println("\n\n------MENU TO CHOSE--------")
		fmt.Print("1.-HTTP GET. \n2.-JSON PARSE \nQ or q to quit\nWrite Code:")
		fmt.Scan(&code)

		switch code {
		case "1":
			http_request_get()
		case "2":
			http_request_get_code_json()
		case "q", "Q":
			fmt.Println("You have chose [",code,"] quit")
		default:
			fmt.Println(code, " is an invalid key")
			time.Sleep(1 * time.Second) //just a timmer to practice
			clear_console()

		}
	}
}
func clear_console(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdin
	cmd.Run()
}


func http_request_get(){
	clear_console()
	fmt.Println("--- HTTP REQUEST GET ---\n")
	
	var url string = "https://api.agify.io?name=michael"
	response, err := http.Get(url)
	
	if err!= nil {
		fmt.Println("error in request :'()")
		return
	}

	defer response.Body.Close() //ensure to close response body

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}


	fmt.Println("body: ", string(body))
	fmt.Println("Status Code:", response.StatusCode )
}

func http_request_get_code_json(){
	clear_console()
	fmt.Println("--- HTTP REQUEST GET ---\n")
	
	var url string = "https://api.agify.io?name=michael"
	response, err := http.Get(url)
	
	if err!= nil {
		fmt.Println("error in request :'(")
		return
	}

	defer response.Body.Close() //ensure to close response body

	var myPerson Persona
	if err:= json.NewDecoder(response.Body).Decode(&myPerson); err != nil{
		panic(err)
	}

	fmt.Println("Name: [", myPerson.Name, "] age: [", myPerson.Age,   "] how many of them exist?: [", myPerson.Count,"]")
}

func post_simulation(){
	data:= map[string]string{
		"usenrbane":"gopher",
		"email":"gopher@example.com"
	}

	jsonData, err := json.Marshal(data)
	if err!= nil{
		panic(err)
	}
	
	resp, err:=http.Post(
		"url",
		"application/json",
		bytes.NewBuffer(jsonData)
	)
	

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("status code: ", resp.StatusCode)
	fmt.Println("Response body: ", string(body))

}


type Persona struct{
	Count int `json: "count"`
	Name string `json: "name"`
	Age int `json:"age"`
}