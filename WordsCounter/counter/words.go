package counter


import (
	"strings"
	"fmt"
)

func CountWords(text []byte)(int){
	fmt.Println("output is ",strings.Fields(string(text)))
	return len(strings.Fields(string(text)))
}