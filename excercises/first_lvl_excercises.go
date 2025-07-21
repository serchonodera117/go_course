package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main(){
	fmt.Print("--- BASIC EXCERCISES ---")

	// first_excercise()
	// map_crud()
	fizz_buzz()
}

func clear_console(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdin
	cmd.Run()
}

// First Excexise 
func first_excercise(){
	
	var arrSize int 
	var code int
	
	fmt.Print("set the size of your array: ")
	fmt.Scan(&arrSize)
	
	var array = make([]int, arrSize)

	fmt.Print("\n\nChose how will be filled \n 1.-normal 0 to N, \n2.-only par, \n\nwrite here: ")	
	fmt.Scan(&code)
	switch code {
		case 1: 
			normal_fill(&array)
		case 2:
			fill_par(&array)
		default:
			fmt.Println("Undefined code")
	}
	
	var mySum int = array_sum(array)

	fmt.Println("array size [",arrSize,"]  array content: ", array )
	fmt.Println("array sum: ", mySum)
}
func normal_fill(arr *[]int){
	for i:=0; i < len(*arr); i++{
		(*arr)[i] = i
	}
}

func fill_par(arr *[]int,){
	for i:=0; i < len(*arr); i++{
		(*arr)[i] = i*2
	}
}
func array_sum(arr []int ) int{
	var x int = 0
	for i,value := range arr {
		fmt.Print(i) // just to try for range
		x+=value
	}
	return x
}



//--------- SECOND EXCERCISE -----------------------------------------
func map_crud(){
	var myMap = make(map[int]string)
	var size int

	fmt.Print("how many objects do you want to introduce? : ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		var myValue string
		fmt.Println("Introduce the value to element [",i,"]")
		fmt.Scan(&myValue)
		myMap[i] = myValue
	}
	
	var code_actions string
	
	for code_actions != "q" && code_actions != "Q"{
		fmt.Print("This is the map: ", myMap, "\n\n\n")
		
		fmt.Println("----ACTIONS----")
		fmt.Print("1.- add \n 2.-delete \n 3.-edit \n Q or q to quit \nwrite here: ")
		fmt.Scan(&code_actions)

		switch code_actions{
			case "1": add_to_map(&myMap)
			case "2": delete_from_map(&myMap)
			case "3": edit_from_map(&myMap)
			case "Q", "q": 
				fmt.Print("you have chosen to get out. :) ")	
			default:
				fmt.Print(code_actions," is an invalid code")	
		}
	}
	
}
func add_to_map(theMap *map[int]string){
	clear_console()
	fmt.Println("----ADD----")
	var myvalue string
	var index int = len(*theMap)

	fmt.Println("ad value to element ", index ,"your map: ")
	fmt.Scan(&myvalue)
	
	(*theMap)[index] = myvalue
}

func delete_from_map(theMap *map[int]string){
	clear_console()
	var key_map int 

	fmt.Println("----DELETE----")
	fmt.Println("this is your map: ", *theMap)
	fmt.Println("Write the key you wnat to delete")
	fmt.Scan(&key_map)

	if _, ok:=(*theMap)[key_map]; !ok {
		fmt.Println("key: ", key_map, " not found, going back. . . ")
		return
	}

	fmt.Println("Deleted: ", key_map, ":", (*theMap)[key_map])
	delete(*theMap, key_map)
}

func edit_from_map(theMap *map[int]string){
	clear_console()
	var key_map int
	var newVal string

	fmt.Println("----EDIT----")
	fmt.Println("this is your map: ", *theMap)

	fmt.Println("chose the key you wnat to edit")
	fmt.Scan(&key_map)

	if _, ok:= (*theMap)[key_map]; !ok {
		fmt.Println("key: ", key_map, " does not exist, returning . . .")
		return
	}

	fmt.Println("key: ", key_map," current value: ", (*theMap)[key_map])
	fmt.Print(key_map, " new value:" )
	fmt.Scan(&newVal)

	(*theMap)[key_map] = newVal

	fmt.Println("succesfully changed")
}


// FIZZ BUZZ
func fizz_buzz(){
	fmt.Println("----FIZZ BUuzz----")
	for i:= 1; i < 11; i++ {
		if( i%2 ==0){
			fmt.Println(i,": fizz")
		}else {
			fmt.Println(i,": buzz")
		}
	}
}