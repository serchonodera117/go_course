package main

import (
	"fmt"
	"time"
)


func main() {
	go task("world") //world execution
	task("hello")
}
func task(s string){
	for i:=0; i<5; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}