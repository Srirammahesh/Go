package main
import (
	"fmt"
	"go_sri/ninjalevel-13/doggy"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  doggy.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(doggy.YearsTwo(20))
}
