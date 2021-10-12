package main
import "fmt"

func main(){
    var n int = 12
    var i int
    for i = 2; i < n; i++ {
        
        if n%i == 0 {
            break
        }
    }
    
    if i == n{
        fmt.Println("Not a prime Number")
    } else{
        fmt.Println("prime Number")
    }
}
