package main
import "fmt"

// Function to right rotate a binary by n bits.
func rrotation( n int, b string) string{
    
    n=n%len(b)
    
    res:= b[len(b)-n:]+b[:len(b)-n]
    return res
}

// Function to right shift a binary by n bits.
func rshift(n int, b string) string{
    
    res:=""
    
    for i:=1 ; i<=n;i++{
        res = res + "0"
    }
    res = res+b
    res = res[:len(b)]
    
    return res
}

// Function to find xor of two binary numbers.
func xor(b1 string, b2 string) string{
    
    res:=""

    for i:=0;i<len(b1);i++{

      if b1[i]==b2[i]{
        res+="0"
      }else{
        res+="1"
      }
    }
    return res
}

// Sigma1 function : Takes a binary(b) and perform xor(right-rotation(17,b),right-rotation(19,b),right-shift(10,b))
func main(){
    
    b := "1011"
    res := xor(xor(rrotation(17,b),rrotation(19,b)),rshift(10,b))

    fmt.Println(res)
}