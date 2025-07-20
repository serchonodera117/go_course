package main

import (
	"fmt"
)

func main (){
	var a int = 10 
	var b int 
	var word_print string = "Hola" + " Mundo"
	

	get_num(&a)
	fmt.Println(word_print, "\n\n Number: ", a, " will be multioied by the number you write: ")
	fmt.Scan(&b)
	multiply_num_by(&a, b)
	fmt.Println("Result is: ", a)

}

func get_num(num *int) {
	*num *= 3 
}

func multiply_num_by(num *int, userNum int){
	*num *= userNum
}