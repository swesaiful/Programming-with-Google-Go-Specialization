package main
import (
    "fmt"
    "strings"
	"bufio"
    "os"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string: ")
    s, _ := reader.ReadString('\n')
	sTrim := strings.TrimSpace(s)    
    if strings.Contains(sTrim,"a") && (sTrim[0]=='i' || sTrim[0]=='I') && (sTrim[len(sTrim)-1]=='n' || sTrim[len(sTrim)-1]=='N'){
        fmt.Print("Found!")
    }else{
        fmt.Print("Not Found!")
    }
}