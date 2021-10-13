package main
import "fmt"

// Function to right shift a binary by n bits.
func main(){
    
    n:=1
    b:="0101"
    res:=""
    
    for i:=1 ; i<=n;i++{
        res = res + "0"
    }
    res = res+b
    res = res[:len(b)]
    
    fmt.Println(res)
}