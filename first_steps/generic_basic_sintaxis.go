package main

import (
	"fmt"
	"reflect"
)

func main(){
	fmt.Printf("---basic sintaxis--- \n\n")
	
	// show_messages()
	data_type_operation()
}


func show_messages(){

	var myString string = "this is my string"
	fmt.Println(myString)

	fmt.Println("type of my data: ", reflect.TypeOf(myString))
}


func data_type_operation(){
	var myNum int = 34
	var myOtherNum float32 = 39.2

	fmt.Println("sum of two variables from diferent data tyope: ", myOtherNum + float32(myNum))
	fmt.Println("type: ", reflect.TypeOf(myOtherNum + float32(myNum)))
}