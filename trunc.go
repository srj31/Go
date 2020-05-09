package main

import "fmt"

func main(){
	fmt.Printf("Type two floating point numbers ")
	var x,y float64

	fmt.Scan(&x, &y)

	trunc := int32(x)
	trunc2 := int32(y)
	fmt.Printf("%d\n%d\n",trunc,trunc2)
}