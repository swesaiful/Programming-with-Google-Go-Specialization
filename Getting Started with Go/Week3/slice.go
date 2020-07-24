package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var userInput string
	var sli []int = make([]int, 3)
	fmt.Println("Slice contains initially --->", sli)
	fmt.Println("Please, enter an integer:")
	for true {
		fmt.Scanln(&userInput)
		if userInput == "X" {
			break
		}
		validInput, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("WRONG INPUT! Please enter an integer or press 'X' to exit ...")
			continue
		}
		sli = append(sli, validInput)
		sort.Ints(sli)
		fmt.Println(sli)
		fmt.Println("Please, enter again an integer or press X to exit ...")
	}
}
