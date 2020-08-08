package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var inputSize = 10

//Swap is a function name
func Swap(numbers []int, i int) {
	indexNum := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = indexNum
}

//BubbleSort is a function name
func BubbleSort(numbers []int) {
	didSwap := true
	for didSwap {
		didSwap = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				didSwap = true
			}
		}
	}
}

// main function
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter integer numbers (up to 10) separated by spaces: ")
	fmt.Println(`Please enter "X" or "x" to exit.`)

	for scanner.Scan() {
		var userInput string = scanner.Text()
		if userInput == "X" || userInput == "x" {
			break
		}

		strNum := strings.Split(userInput, " ")
		var intSli = make([]int, 0)
		for _, i := range strNum {
			intNum, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("WRONG INPUT!")
				fmt.Printf("Please enter integer numbers (up to 10) separated by spaces: ")
				continue
			}
			intSli = append(intSli, intNum)
		}
		fmt.Printf("len(intSlice): %d; cap(intSlice): %d\n", len(intSli), cap(intSli))
		if len(intSli) > inputSize {
			fmt.Printf("Please enter a maximum of %d integers\n", inputSize)
			fmt.Println("Please enter integer numbers (up to 10) separated by spaces: ")
			continue
		}
		BubbleSort(intSli)
		fmt.Printf("Integer numbers in sorted order: %v\n", intSli)
		fmt.Println("\nPlease enter integer numbers (up to 10) separated by spaces: ")
	}
}
