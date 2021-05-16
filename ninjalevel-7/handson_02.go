package main
import "fmt"
type person struct{
	first string
	last string
	age int
}

func changeme(s *person){
	(*s).first="mahi"//(*value).anything
	s.last="bhai"//value.anything (both are possible)
	s.age=39
}
func main(){
	p:=person{
		"dhoni",
		"baiya",
		37,
	}
	fmt.Println(p)
	changeme(&p)
	fmt.Println(p)
}