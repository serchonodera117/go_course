package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("excercises")

	// invert_a_given_string("esteganografía")
	// sort_an_array([]int{34,4,1,7,2,6,10})

// 	var mynum int = 5
// 	fmt.Print("num: ", mynum)
// 	factorial(&mynum)
// 	fmt.Print(" factorial ", mynum)
	fibonacchi(10)
}


func invert_a_given_string(word string) {
	var reversed_word string
	var stringSplit = strings.Split(word, "")

	for i:=len(stringSplit)-1; i > -1; i-- {
		 reversed_word += stringSplit[i]
	}


	fmt.Printf("original word: %s, reverted word: %s", word, reversed_word)
}

func sort_an_array(givenArray []int){
	fmt.Println("original Array: ", givenArray)

	for i:=0; i < len(givenArray)-1; i++{
		for j:=i+1; j < len(givenArray); j++{
			if(givenArray[i] < givenArray[j]){
				currentPos := givenArray[i]
				givenArray[i] = givenArray[j]
				givenArray[j] = currentPos
			}
		}
	}

	fmt.Println("sorted Array: ", givenArray)
}


func factorial(n *int){
	var result int = 1
	
	for i:= 1;  i <= *n; i++{
		result *= i 
	}
	*n = result

}

func fibonacchi(n int){	
	var a int = 0
	var b int = 1

	for i:=1; i <= n; i++ {		
		var c = a + b
		a = b
		b = c
		
		fmt.Printf("Index: %d, Value: %d  \n", i,b)
	}
}