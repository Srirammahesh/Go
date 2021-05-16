package main
import "fmt"
type person struct{
	name string
	age int
}
func (p *person) speak(){
	fmt.Println("my name is",p.name)
}
type human interface{
	speak()
}
func saysomething(h human){
	h.speak()
}

func main(){
	p:=person{
		"sriram",21,
	}
	//k:=&p
	//k.speak() this works
	p.speak()// this too works cause although pointer needed for person it executes
	//saysomething(p)// but this wont work cause speak need pointer to person and person implicitly inmplements human means pointer has to be passed for the speak function
	saysomething(&p)
}