package main
import "fmt"
func main() {
    var f float64
    fmt.Printf("Enter a floating-point number: ")
    fmt.Scanf("%f", &f)
    i := int64(f)
    fmt.Printf("The truncated version of %g is: %d\n",f, i)
}