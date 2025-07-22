package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	fmt.Println("---LOOPS----")

	var code string
	
	for code != "q" && code != "Q" {
		fmt.Println("\n\n------MENU TO CHOSE--------")
		fmt.Print("1.-for. \n2.-while.\n3.-for range \n4.-for range with slices \nQ or q to quit\nWrite Code:")
		fmt.Scan(&code)

		switch code {
		case "1":
			for_loop()
		case "2":
			while_loop()
		case "3":
			for_range()
		case "4":
			for_range_with_slices()
		case "q", "Q":
			fmt.Println("You have chose [",code,"] quit")
		default:
			fmt.Println(code, " is an invalid key")
			time.Sleep(1 * time.Second) //just a timmer to practice
			clear_console()

		}
	}

	fmt.Print("Have a good day :)")


}

func for_loop() {
	clear_console()
	fmt.Println("---FOR LOOP---")
	var mySlice []int = []int{}

	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, i)
	}

	fmt.Println(mySlice)
}

func while_loop() {
	clear_console()
	fmt.Println("---WHILE LOOP---")

	var limit int16 = 0

	for limit < 10 {
		fmt.Println(limit)
		limit++
	}
}

func for_range() {
	clear_console()
	fmt.Println("---FOR RANGE--- \n use with maps")
	var myMap = make(map[string]int)
	
	for i:=0; i < 10; i++ {
		myMap[strconv.Itoa(i)] = i
	}
	fmt.Println("| Index  |   Value |")
	for index, value := range myMap{	
		fmt.Println("Index: [", index, "]:", value)
	}
}


//endless for  "for { code }"
func for_range_with_slices(){
	clear_console()
	fmt.Println("---FOR RANGE--- \n use with slices")
	var mySlice []int = []int{}
	
	for i:=0; i < 10; i++ {
		mySlice = append(mySlice, i)
	}
	fmt.Println("| Index  |   Value |")
	for index, value := range mySlice{	
		fmt.Println("Index: [", index, "]:", value)
	}
}
func clear_console(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdin
	cmd.Run()
}
