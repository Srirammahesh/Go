package main
import "fmt"
func main(){
	i:=[]int{1,2,3,4,5,6,7,8,9}
	k:=sum(i...)
	fmt.Println(k())
	//or use this
	//k:=sum(i...)()
	//fmt.Println(k)
}
func sum(x ...int) func() int{
	return func() int{
		total:=0
		for _,v :=range x{
			total+=v
		}
		return total
	}
}
