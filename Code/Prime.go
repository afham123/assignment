package main
import "fmt"

// function to check prime number.
func main(){
    
    n:=16
    result:=true
    for i:=2; i<n;i++{
        if n%i==0{
            result=false
        }
    }
    fmt.Println(result)
}