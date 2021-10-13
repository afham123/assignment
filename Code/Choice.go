package main
import "fmt"
//import "strconv"

// Function to read each bits of first binary and add second to result if it is one else add third to result.
func main(){
    
    b1:="1011"
    b2:="1010"
    b3:="0011"
    res:=""

    for i:=0;i<len(b1);i++{
      a := string(b1[i])
      if (a=="1"){
        res=res+ string(b2[i])
      }else{
        res = res + string(b3[i])
      }
    }
    fmt.Println(res)
}