package main
import "fmt"

// Function to find xor of two binary numbers.
func main(){
    
    b1:="0011"
    b2:="1010"
    res:=""

    for i:=0;i<len(b1);i++{

      if b1[i]==b2[i]{
        res+="0"
      }else{
        res+="1"
      }
    }
    fmt.Println(res)
}