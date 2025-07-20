package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main(){
	fmt.Printf("---basic sintaxis--- \n\n")
	
	// show_messages()
	// data_type_operation()
	// arrays_data_strcture()
	// map_sintaxis()
	// list_sintaxis()

	fmt.Println("---stuct ---")
	var my_instance MyStruct = MyStruct{"serch",24}

	my_instance.print_propeties()
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

func arrays_data_strcture(){
	fmt.Println("---  arrays ---")

	//basic declaration
	var myArray [3] int = [3]int{1, 2, 3}
	var myStArray [5] string
	var mySlice = []int16{1,2,3}

	myStArray[0] = "cero"
	myStArray[1] = "uno"
	myStArray[2] = "dos"
	myStArray[3] = "tres"
	myStArray[4] = "cuatro"

	fmt.Println("int array: ", myArray)
	fmt.Println("string array:  ", myStArray)
	
	fmt.Print("\n this is a type fo dynamic array called slice")
	
	fmt.Println("my original slice:  ", mySlice)
	
	mySlice = append(mySlice, 5)
	fmt.Println("my original plus a new value:  ", mySlice)
	
	mySlice = append(mySlice, []int16{12,3,4,4}...)
	fmt.Println("my original plus a new slice:  ", mySlice)
}

func map_sintaxis(){
	fmt.Println("--- maps ---")

	var myMap  = make(map[string]int) //map empty
	var key_to_look_for string
	fmt.Println("empty map", myMap, "\n")

	myMap["a"] = 1
	myMap["b"] = 2
	myMap["c"] = 3
	myMap["d"] = 4
	myMap["e"] = 5    //map no longer empty

	fmt.Println("map with values", myMap)
	fmt.Println("map value a: ", myMap["a"])
	fmt.Println("map value b: ", myMap["b"])
	fmt.Println("map value non exists : ", myMap["z"])
	
	fmt.Print("look for a key in the map: ")
	fmt.Scan(&key_to_look_for)

	if(myMap[key_to_look_for]==0) {
		fmt.Println(key_to_look_for, " was not found")
		// return
	}else {
		fmt.Println("value of [ ",key_to_look_for," ] is:", myMap[key_to_look_for])
	}

	var map_by_structure = map[string]int{"a":10,"b":30,"c":2}

	fmt.Println("map by crating its structure: ", map_by_structure)
}


func list_sintaxis(){
	fmt.Println("--- lists ---")
	var myList = list.New() //lists also works as a set
	fmt.Println("empty list :  ", myList) //is necesary to use value to acces to the values of the list

	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)

	fmt.Println(myList.Back().Value)

}

//struct (similar to classes but they are not classes)
type MyStruct struct {
	name string
	age int
}
func (p MyStruct)print_propeties(){
	fmt.Print("name: ", p.name, "\n age: ", p.age)
}