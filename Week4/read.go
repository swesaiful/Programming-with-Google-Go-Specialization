package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name is a type name
type Name struct {
	fname string
	lname string
}

func main() {

	readInput := bufio.NewScanner(os.Stdin)
	fmt.Println("File Path: ")
	readInput.Scan()
	fPath := readInput.Text()

	f, err := os.Open(fPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sli := make([]Name, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ns := strings.Fields(scanner.Text())
		p := Name{
			fname: ns[0],
			lname: ns[1],
		}
		sli = append(sli, p)
	}
	fmt.Println("Iterate name strings through the slice of structs ...")
	for i := 0; i < len(sli); i++ {
		fmt.Printf("fname: [%-10s]; lname: [%-10s]\n", sli[i].fname, sli[i].lname)
	}

}
