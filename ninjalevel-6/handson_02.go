package main
import "fmt"
func main(){
	l:=[]int{1,2,3,4,5}
	f:=foo(l ...)
	fmt.Println(f)
	fmt.Println(bar(l))

}
func foo(x ...int) int{
	sum:=0
	for _,v:=range x{
		sum+=v
	}
	return sum
}

func bar(x []int)int{
	sum:=0
	for _,v:=range x{
		sum+=v
	}
	return sum
}
