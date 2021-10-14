package main
import (
	"bufio"
	//"bytes"
	"fmt"
	//"math"
	"os"
	"crypto/sha256"
	"encoding/base64"
)

var keyValue map[string]string
var keyHash map[string]string

//Function that takes key, value pair as an argument and 
//store it in key,value in keyValue hash map key hash in keyHash hash map 
func Put(key string, value string){
    
    h := sha256.New()
    h.Write([]byte(value))
    
    sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
    
    fmt.Println(sha)
    keyValue[key] = value //storing key value in keyValue hash map
    keyHash[key] = sha  // storing key hash in keyHash hash map 
}



func main() {
  
  
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Enter the number of key value pair to store in kv store")
  num := readNum(reader)
  
  for i:=1;i<=num;i++{
        
        key := readString(reader)
        val := readString(reader)
        
        Put(key,val)
    }
    
    fmt.Println(keyValue)
    fmt.Println(keyHash)
}

// Reading input functions
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}