package main

import (
	"fmt"
	"strings"
)

func main(){

	var str string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&str)
	str = strings.ToLower(str)
	if strings.Contains(str,"i") && strings.Contains(str, "a") && strings.Contains(str,"n") {
		fmt.Printf("Found!\n")
	} else{
		fmt.Printf("Not Found!\n")
	}
}