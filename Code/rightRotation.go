package main
import "fmt"

// Function to right rotate a binary by n bits.
func main(){
    
    b:="0011"
    n:=2
    n = n%len(b)
    
    res:= b[len(b)-n:]+b[:len(b)-n]
    fmt.Println(res)
}