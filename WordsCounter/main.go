package main

import (
	"wordscounter/counter"
	"log"
	"os"
	"fmt"
)

var filename = "words/go.txt"


func main(){

   if len(os.Args)>1{
	filename = os.Args[0]
   }


  
   text,err := os.ReadFile(filename)
   if err!=nil{
	log.Fatal("Error while reading the file",filename)
   }

   fmt.Println("The total words in this file is",counter.CountWords(text))
	
}