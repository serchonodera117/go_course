package main

import "fmt"



type Performance interface {
	Attack() string
	Receive_damage() int
	Heal() int
	Defeated() string
}

type Enemy struct {
	en_name string
	life int
	attack int
	defense int
}

type Boss struct {
	en_name string
	life int
	attack int
	defense int
	animation string
	combo string
}

func (e Enemy) Attack() string{
	var message string = e.en_name + " Is attacking"
	// fmt.Println(message)
	return message
}

func (e Boss) Attack() string{
	var message string  = e.en_name  + " bos is attacking with " + e.animation  + "animation"
	// fmt.Println(message)
	return message
}
func (e Boss) Receive_damage() int {
	return 10	
}

func (e Boss) Heal() int{
	return 20
}

func (e Boss) Defeated() string {
	return e.en_name + " has been defeated"
}

//func that allows an interfaz 'shape'
func performance_info(e Performance){
	fmt.Print("attack: ", e.Attack())
	fmt.Print(e.Defeated())
}

func main(){
	var bos_a Boss = Boss{
		en_name: "bowser",
		life: 100,
		attack: 50,
		defense: 70,
		animation: "ulti_anim",
		combo: "multiple attaks",
	}

	performance_info(bos_a)
	// fmt.Print(bos_a.Attack())
}