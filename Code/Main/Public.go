package main
import (
	"bufio"
	//"bytes"
	"fmt"
	//"math"
	"os"
	"crypto/sha256"
	"encoding/base64"
	"encoding/csv"
	"log"
)

// struct to store keyvalue and hash file.
type Data struct {
    val string
    hash string
}

//struct to store userid password
type authentication struct {
    password string
}

var keyValue = map[string]Data{}  // Hash map to store key(dates in string) and its news content(in string type) and its hash value(in string type)
var authentication = map[string]authentication{}  // Hash map to store userid password.


//Function that takes key, value pair as an argument and 
//store it in key,value and its hash256 in keyValue hash map. 
func Put(key string, value string){
    
    h := sha256.New()
    h.Write([]byte(value))
    sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
    
    //fmt.Println(sha)
    var n Data 
    
    n.val = value //storing key value in keyValue hash map.
    n.hash = sha  // storing key hash in keyValue hash map.
    
    keyValue[key] = n
    
    save()   // Saving the changes.
}

//Function to readin CSV file.
func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }
    return records
}

//Function to read local data.
func read(){
    
    data = readCsvFile()
    
    
    
    
    
}



// Function to return news contnet and its Hash Value
func Get(key string) string{
    
    if Fingerprint(key,keyValue[key]){
        return keyValue[key]
    }else{
        Delete(key)
        update()
        return 
    }
}


//Update with new value
func update(key string, val string){
    
    h := sha256.New()
    h.Write([]byte(value))
    sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
    
    n := keyValue[key]
    n.val = val
    n.hash = sha
    keyValue[key] = n
}

//Function to check the fingerprint and return a boolean result.
// That is true if mached and false if does not matched.
func Fingerprint(key string, val string) bool{
    
    h := sha256.New()
    h.Write([]byte(value))
    sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
    
    if keyValue[key].hash == sha { // will be false if hash value of original content matches the hash value of news content.
        return true
    }
    return false
}

// Delete the given key from the store
func Delete(key string){
    n := keyValue[key]
    n.val = ""
    n.hash = ""
    keyValue[key] = n
}

// function to save the content locally in csv.
func save(){
    
    file, _ := os.Create("export.csv")
    defer file.Close()
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // define column headers
    headers := []string{
        "Date(key)",
        "Content",
        "Hash",
    }
    
    // write column headers
    writer.Write(headers)
    
    for key := range keyValue {
    
        r := make([]string, 0, 1+len(headers)) // capacity of 4, 1 + the number of properties your struct has & the number of column headers you are passing
        r = append(
            r,
            key,
            keyValue[key].val,
            keyValue[key].hash,
        )
        writer.Write(r)
    }
}

//Function for userid password verification 
func check( usid string, pass string)bool{
    
    if authentication[usid] { // will be false if userid is not in the map
        if authentication[usid].password == pass{
            return true
        }
        return false
    }
    return false
}


// Driver Function
func main() {
  
  
    authentication["arijit.Das"].password = "UITBU.Arijit" // Adding user id and password
    // Reading the local csv data.
    
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Println("Switch Case")
    fmt.Println("Enter 1 to add new Data")
    fmt.Println("Enter 2 to read news")
    fmt.Println("Enter 3 to exit")
    num := readNum(reader)   
    
    for num !=3{
        switch num {
        case 1:
            fmt.Printf("Enter user : ")
            fmt.Printf("Enter password : ")
            
            usid := readString(reader)
            pass := readString(reader)
            
            if check(usid,pas)
            {
                fmt.Println("Enter the key value pair")
                key := readString(reader)
                val := readString(reader)
                Put(key,val)    // saving key value pair after authentication.
            }
            break
        case 2:
            fmt.Println("Enter the date you want to read news")
            key := readString(reader)
            fmt.Println(Get(key))
            break
        case 3:
            fmt.Println("Thanks for using this application, Have a nice day")
            break
        default:
            fmt.Println("Wrong choice Entered")
        }
        if num == 3{
            break            // Breaking the switch case if the user enters 3
        }else{
            num = readNum(reader)
        }
    }
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
