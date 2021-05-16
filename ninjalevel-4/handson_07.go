package main
import "fmt"
func main(){
	x:=[]string{"james","bond","shaken,not stirred"}
	y:=[]string{"miss","moneypenny","hellojames"}
	z:=[][]string{x,y}
	fmt.Println(z)
	for _,k:=range z{
		for _,v:=range k{
			fmt.Println(v)
		}
	}
}
