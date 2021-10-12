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
    fmt.Printf(res)
}
