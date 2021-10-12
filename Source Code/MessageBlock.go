package main
import "fmt"
import "strconv"

//function to convert decimal to binary 
func binary(n int) string{
    var res = ""
    for (n>0){
        r := (n%2)
        res=strconv.Itoa(r)+res
        n/=2
    }
    return res
}

// function to convert text to binary 
func textToBinary(text string) string{
    res := ""
    
    for i:=0; i<len(text); i++{
        res = res + binary(int(text[i]))
    }
    
    res = res + binary(len(text))
    return res
}

// function to break the text into a blocks of 512 bits each.
func break_512(st string) map[int]string{
    
    res := make(map[int]string)
    
    for i:=1; i<=len(st)/512; i++{
        
        res[i] = st[(i-1)*512:i*512]
    }
        
    return res
}

// function to padd zeros in the begining and length of binary text at the end of the binary text.
func main(){
    text := "My name is Afham Fardeen"
    res := textToBinary(text) + binary(len(text))
    
    t:= ""
    for i:=0; i<(512-(len(res)%512)) ; i++{
        t = "0" + t
    }
    res = t+res
    
    result := break_512(res)
    fmt.Println(result)
}