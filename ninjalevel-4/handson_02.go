package main

import "fmt"

func main(){
	var x = []int{0,1,2,3,4,5,6,7,8,9}
	for _,v :=range x{
		fmt.Println(v)
	}
	fmt.Printf("%T",x)
}
