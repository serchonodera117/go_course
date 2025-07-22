package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("excercises")

	invert_a_given_string("esteganografía")
}


func invert_a_given_string(word string) {
	var reversed_word string
	var stringSplit = strings.Split(word, "")

	for i:=len(stringSplit)-1; i > -1; i-- {
		 reversed_word += stringSplit[i]
	}


	fmt.Printf("original word: %s, reverted word: %s", word, reversed_word)
}