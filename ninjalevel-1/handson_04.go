package main

import "fmt"

type sri int
var x sri
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n",x)
	x=42
	fmt.Println(x)
}
