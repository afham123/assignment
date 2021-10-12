package main
import "fmt"
import "math"
import "strconv"
//import "reflect"

// Convert Decimal to Binary
func binary(n int)string{
    var res = ""
    for (n>0){
        r := (n%2)
        res=strconv.Itoa(r)+res
        n/=2
    }
    return res
}

// function to check prime number.
func prime(n int)bool{
    
    result:=true
    for i:=2; i<n;i++{
        if n%i==0{
            result=false
        }
    }
    return result
}

// Function to create constant(k[64]) for hashing
func main(){
    i:=0
    a:=2
    var lst[64] int
    var res[64] string
    
    for i<64{
        if prime(a){
            lst[i] = a
            i+=1
        }
        a+=1
    }
    fmt.Println(lst)
    
    for i:=0; i<64; i++{
        a := float64(math.Pow(float64(lst[i]),0.3333))
        a = (a - float64(int(a))) * math.Pow(float64(2),32)
        b := (int)(a)
        res[i]= binary(b)
    }

    for i:=0; i<64;i++{
        if len(res[i])!=32{
            t:=""
            for j:=0;j< 32-len(res[i]);j++{
                t = t+"0"
            }
            res[i] = res[i]+t
        }
    }
    fmt.Println(res)
}