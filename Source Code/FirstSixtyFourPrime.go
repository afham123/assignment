package main
import "fmt"

// function to check prime number.
func prime(n int)bool{
    
    result:=true
    for i:=2; i<n;i++{
        if n%i==0{
            result=false
        }
    }
    return result
}

// Function to get first 64 prime number.
func main(){
    i:=0
    a:=2
    var lst[64] int
    
    for i<64{
        if prime(a){
            lst[i] = a
            i+=1
        }
        a+=1
    }
    fmt.Println(lst)
}