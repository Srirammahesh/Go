package main
import "fmt"
func main(){
	i:=[]int{1,2,3,4,5,6,7,8,9}
	f:=oddsum(total,i)
	fmt.Println(f)


}
func total(x []int) int {
	tot:=0
	for _,v :=range x{
		tot+=v
	}
	return tot
}

func oddsum(l func(x []int) int,k []int) int{
	var arr []int
	for _,v :=range k{
		if v%2==0{
			arr=append(arr,v)
		}
	}
	return l(arr)
}