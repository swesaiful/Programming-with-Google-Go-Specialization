package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)

func main() {

	readInput := bufio.NewScanner(os.Stdin)

	fmt.Print("Username: ")
	readInput.Scan()
	name := readInput.Text()

	fmt.Print("Address: ")
	readInput.Scan()
	address := readInput.Text()
	//fmt.Printf("Username: %s, Address: %s\n", name, addr)

	person := make(map[string]string)
	person["name"] = name
	person["address"] = address

	jsonperson, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nJSON: %s\n", jsonperson)
}
