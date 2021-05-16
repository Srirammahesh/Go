package main

import "fmt"

type square struct{
	len float64
}
type circle struct{
	radius float64
}
func (s square) area() float64{
	return s.len*s.len
}
func (c circle) area() float64{
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64

}
func info(s shape){
	fmt.Println(s.area())
}
func main(){
	s:=square{
		len:5,
	}
	c:=circle{
		radius:7,
	}
	info(s)
	info(c)
}