package main

import (
	"fmt"
	"strings"
)

// Animal is a type name and struct contains three fields
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Global variable declaration

var aMap = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}
var aName = []string{"cow", "bird", "snake"}
var aAction = []string{"eat", "move", "speak"}

// Eat method will invoke Animal
func (a Animal) Eat() {
	fmt.Printf("%s\n", a.food)
}

// Move method will invoke Animal
func (a Animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

// Speak method will invoke Animal
func (a Animal) Speak() {
	fmt.Printf("%s\n", a.noise)
}

// userPrompt is a function name
func userPrompt() {
	fmt.Println("Please enter Animal name and action [space separated].")
	fmt.Printf("Type Animal name from this list --> (%s)\n", strings.Join(aName, ","))
	fmt.Printf("Type action from this list ---> (%s)\n", strings.Join(aAction, ","))
}


func main() {
	userPrompt()
	for {
		fmt.Print(">")
		var animal, action string = "", ""
		fmt.Scan(&animal, &action)
		
		switch {
		case action == "eat":
			aMap[animal].Eat()
		case action == "move":
			aMap[animal].Move()
		case action == "speak":
			aMap[animal].Speak()
		default:
			fmt.Println("WRONG INPUT! TRY AGAIN ...")
			break
		}

	}
}
