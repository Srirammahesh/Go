package main
import "fmt"
type person struct{
	first string
	last string
	icecream []string
}
func main(){
	p1:=person{
		first:"sri",
		last:"spark",
		icecream:[]string{
			"butterscotch","vanilla",
		},
	}
	p2:=person{
		first:"iron",
		last:"man",
		icecream:[]string{
			"blackcurrent","mango",
		},
	}
	fmt.Println(p1,p2)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range (p1.icecream){
		fmt.Println(i,v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range (p2.icecream){
		fmt.Println(i,v)
	}


}
