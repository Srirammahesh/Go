package main

import "fmt"

var a=42
func main(){
	fmt.Printf("%d \t %b \t %#x",a,a,a)
	b:=a<<1
	fmt.Printf("%d \t %b \t %#x",b,b,b)
}
