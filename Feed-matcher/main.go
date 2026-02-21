package main 

import (
	"log"
	"os"
	"fmt"
	"feedmatcher/search"
	_"feedmatcher/matchers"
)


func init(){
	fmt.Println("------Initialzed")
	log.SetOutput(os.Stdout) // logs to be printed in stdout 
}

func main(){
	var docToSearch string 
	fmt.Scan(&docToSearch)
	search.Run(docToSearch)
}