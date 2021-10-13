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

//Function to take text and convert it into binary add padd zeros in the begining and add length of 
//text at the end of the binary and create Message Blocks by dividing it into a blocks of 512 bits.
func messageBlocks(text string) map[int]string{

    res := textToBinary(text) + binary(len(text))
    
    t:= ""
    for i:=0; i<(512-(len(res)%512)) ; i++{
        t = "0" + t
    }
    res = t+res
    
    result := break_512(res)
    return result
}

// Driver function
func main(){
    
    res := "Main function"
    fmt.Println(res)
    
    text := "My name is Afham Fardeen. I am Pursuing Bachelor of Engineering and studing in The University Institute of Technology, Burdwan University."
    fmt.Println("Original Text : "+text)
    
    fmt.Println("Message Blocks")
    fmt.Println(messageBlocks(text))
}
