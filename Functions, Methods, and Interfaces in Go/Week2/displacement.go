package main

import (
	"math"
	"bufio"
	"os"
	"log"
	"strconv"
	"fmt"
)

// GenDisplaceFn is a function name
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*(math.Pow(t, 2)*0.5) + (v0 * t) + (s0))
	}
	return fn
}

// readUserInput is a function name
func readUserInput(inputType string) float64 {

	var inputValue float64

	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	strValue := inputType
	validInput := true
	for validInput {
		validInput = false
		fmt.Printf("%s or (X to exit): ", inputType)
		scanner.Scan()
		strValue = scanner.Text()
		if strValue == "X" {
			os.Exit(0)
		}
		
		inputValue, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			fmt.Println("Invalid Input!")
		} else {
			validInput = true
			return inputValue
		}

	}
	return inputValue
}

// main function
func main(){
	var acceleration, initialVelocity, initialDisplacement, time float64

    acceleration = readUserInput("accelleration")
    initialVelocity = readUserInput("velocity")
    initialDisplacement = readUserInput("displacement")
    time = readUserInput("time")

	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)
	
	fmt.Printf("\nacceleration: %f\n", acceleration)
	fmt.Printf("velocity:     %f\n", initialVelocity)
	fmt.Printf("displacement: %f\n", initialDisplacement)
	fmt.Printf("time:         %f\n", time)
    fmt.Printf("Displacement: %f\n", fn(time))

}