//function to get binary of 20 bits of a number
package main
import "fmt"
import "strconv"

func main(){
    n := 12
    
    var res = ""
    for (n>0){
        r := (n%2)
        res=strconv.Itoa(r)+res
        n/=2
    }
    t:=""
    for i:=0; i < 20-len(res) ;i++{
        t=t+"0"
    }
    
    res = t + res
    fmt.Printf(res)
}