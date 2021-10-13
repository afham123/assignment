package main
import "fmt"

// Function to get the majority of bits at that place.
func main(){
    
    b1:="0011"
    b2:="1010"
    b3:="0011"
    res:=""

    for i:=0;i<len(b1);i++{
      a := string(b1[i])+string(b2[i])+string(b3[i])
      if (a=="101"|| a=="110" || a=="011"  || a =="111"){
        res=res+"1"
      }else{
        res = res + "0"
      }
    }
    fmt.Println(res)
}