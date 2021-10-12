package main
import "fmt"

func main(){
    n := int64(255)
    h := fmt.Sprintf("%x", n)
    fmt.Printf("%s", h)
}
