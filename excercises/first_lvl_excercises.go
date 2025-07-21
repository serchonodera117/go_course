package main

import (
	"fmt"
)

func main(){
	fmt.Print("--- BASIC EXCERCISES ---")

	first_excercise()
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