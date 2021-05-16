package main
import "fmt"
func main(){
	p:=struct{
		first string
		last string
		age int
		favfood []string
		friends map[string]int
	}{
		first:"sri",
		last:"spark",
		age:22,
		favfood:[]string{"pizza","burger"},
		friends: map[string]int{
			"sri":111,
			"spark":001,
		},
	}
	fmt.Println(p.first)
	fmt.Println(p.last)
	fmt.Println(p.age)
	for i,v :=range p.favfood{
		fmt.Println(i,v)
	}
	for i,v :=range p.friends{
		fmt.Println(i,v)
	}

}
