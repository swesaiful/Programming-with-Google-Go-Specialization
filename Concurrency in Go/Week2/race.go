/*
* A race condition is when the outcome of the program does not depend on the interleaving. 
* Because interleaving is non-deterministic, and it needs to be avoided.

* Race conditions occur due to communication between Goroutines, and communication is done 
* through a shared variable.

*/

package main

import (
	"fmt"
	"time"
)

func race(str string) {
	fmt.Println(str)
}

func main() {
	for i := 0; i < 3; i++ {

		// The results of the following strings "Go" and "Gopher" are uncertain 
		// due to the uncertainty of Goroutine scheduling mechanism, 
		// and this is where the race condition occurs.
		// The display of the string "Go" and "Gopher" is different each time 
		// the program is run due to the race condition.

	go race("Go")
	go race("Gopher")
	time.Sleep(1 * time.Second)
	}
}
