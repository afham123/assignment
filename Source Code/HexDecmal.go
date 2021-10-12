import "fmt"
import "strconv"

func main(){
    n := 280
    
    var res = ""
    for (n>0){
        r := (n%16)
        if r > 9{
            res=  string(87 + r) + res 
        }else{
            res=strconv.Itoa(r)+res
        }
        n/=16
    }
    fmt.Printf(res)
}
