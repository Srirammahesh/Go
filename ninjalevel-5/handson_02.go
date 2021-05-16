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
	m:=map[string]person{
		p1.last:p1,
		p2.last:p2,
	}
	for i,v:=range m{
		fmt.Println(i)
		fmt.Println("\t", v.first)
		fmt.Println("\t ",v.last)
		for k,x :=range v.icecream{
			fmt.Println("\t",k,x)
		}
		}
	}
