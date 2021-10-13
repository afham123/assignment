package main
import (
  "fmt"
	"strconv"
	"strings"
)

// Function to find add two binary numbers.
func add(a string, b string) string {
	if len(b) > len(a) {
		a, b = b, a
	}

	res := make([]string, len(a)+1)
	i, j, k, c := len(a)-1, len(b)-1, len(a), 0
	for i >= 0 && j >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		bj, _ := strconv.Atoi(string(b[j]))
		res[k] = strconv.Itoa((ai + bj + c) % 2)
		c = (ai + bj + c) / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		ai, _ := strconv.Atoi(string(a[i]))
		res[k] = strconv.Itoa((ai + c) % 2)
		c = (ai + c) / 2
		i--
		k--
	}
	if c > 0 {
		res[k] = strconv.Itoa(c)
	}
	return strings.Join(res, "")
}

// Driver function
func main(){
    
    b1 := "1011"
    b2 := "0011"
    res := add(b1,b2)

    fmt.Println(res)
}
