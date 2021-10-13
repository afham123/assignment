package main

import (
    "crypto/sha256"
    "fmt"
)

type KVStore interface {
    
	// Given a key, retrieve the corresponding value from the store.
	Get(key string) (value []byte, exists bool)
	
	// Save the given value for the given key into store.
	Put(key string, value []byte)
	
	// Delete the given key from the store.
	Delete(key string)
	
	// Calculate the fingerprint for the store. Any two stores which have the same data, should have the same fingerprint, and vice-versa.
	Fingerprint() []byte
}
func Get(key string) (value []byte, exists bool) {
    
    exit := true
    var val [5]byte
    
    return val,exit
}
func Put(key string, value []byte){
    
    
}
func Delete(key string){
    

}
func Fingerprint() []byte{
    
    var val [5]byte
    
    return val
}


func main() {
    
    text = "This is Afham fardeen"
    text1 = text1+"m"
    sh := sha256.Sum256([]byte(text))
    sh1 := sha256.Sum256([]byte(text1))
    fmt.Printf("%x\n", sh)
    fmt.Printf("%x", sh1)
    
    
}
