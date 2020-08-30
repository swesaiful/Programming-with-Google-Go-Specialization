package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Animal interface
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

// Cow is a struct type with three fields
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Eat method will invoke Cow
func (a Cow) Eat() {
	fmt.Printf("%s\n", a.food)
}

// Move method will invoke Cow
func (a Cow) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

// Speak method will invoke Cow
func (a Cow) Speak() {
	fmt.Printf("%s\n", a.noise)
}

// Bird is a struct type with three fields
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Eat method will invoke Bird
func (a Bird) Eat() {
	fmt.Printf("%s\n", a.food)
}

// Move method will invoke Bird
func (a Bird) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

// Speak method will invoke Bird
func (a Bird) Speak() {
	fmt.Printf("%s\n", a.noise)
}

// Snake is a struct type with three fields
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Eat method will invoke Snake
func (a Snake) Eat() {
	fmt.Printf("%s\n", a.food)
}

// Move method will invoke Snake
func (a Snake) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

// Speak method will invoke Snake
func (a Snake) Speak() {
	fmt.Printf("%s\n", a.noise)
}

func main() {

	var aMap map[string]string
	aMap = make(map[string]string)


	for {
		var command, aName, aType string
		fmt.Println(`Request for type "newanimal" or "query" command: `)
		fmt.Printf("\n>")

		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
		scanner.Scan()
		strValue := scanner.Text()
		validInput := strings.Fields(strValue)

		if len(validInput) == 3 {

			command = validInput[0]
			aName = validInput[1]
			aType = validInput[2]

		} else {
			fmt.Println(`Input "command" is not equal to 3!`)
			continue
		}

		switch command {
		case "newanimal":

			_, ok := aMap[aName]
			if ok != true {
				aMap[aName] = aType
				fmt.Println("Created it!")
			} else {
				fmt.Printf("%s: Animal type is Invalid\n", aType)
			}
		case "query":

			switch aMap[aName] {
			case "cow":
				cow := Cow{"grass", "walk", "moo"}
				switch aType {
				case "eat":
					cow.Eat()
				case "move":
					cow.Move()
				case "speak":
					cow.Speak()
				default:
					fmt.Printf("%s: Invalid Request\n", aType)
					continue
				}
			case "bird":
				bird := Bird{"worms", "fly", "peep"}
				switch aType {
				case "eat":
					bird.Eat()
				case "move":
					bird.Move()
				case "speak":
					bird.Speak()
				default:
					fmt.Printf("%s: Invalid Request\n", aType)
					continue
				}
			case "snake":
				snake := Snake{"mice", "slither", "hsss"}
				switch aType {
				case "eat":
					snake.Eat()
				case "move":
					snake.Move()
				case "speak":
					snake.Speak()
				default:
					fmt.Printf("%s: Invalid Request\n", aType)
					continue
				}
			default:
				fmt.Printf("%s: Invalid Request\n", aMap[aName])
				continue
			}
		default:
			fmt.Printf("%s: Invalid Request\n", command)
			continue
		}
	}
}