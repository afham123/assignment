//function to convert a text to binary
package main
import "fmt"
import "strconv"

func binary(n int) string{
    var res = ""
    for (n>0){
        r := (n%2)
        res=strconv.Itoa(r)+res
        n/=2
    }
    return res
}


func main(){
    text := "My name is Afham Fardeen"
    res := ""
    
    for i:=0; i<len(text); i++{
        res = res + binary(int(text[i]))
    }
    
    res = res + binary(len(text))
    fmt.Println(res)
    
}