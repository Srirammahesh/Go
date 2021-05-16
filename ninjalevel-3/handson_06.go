package main

import "fmt"

func main(){
	a:=5
	for i:=0;i<10;i++{
		if i==a{
			break
		}
		fmt.Println(i)
	}
}
