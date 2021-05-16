package main
import "fmt"
func main(){
	defer foo()
	fmt.Println("prints in main")
}
func foo(){
	defer func(){
		fmt.Println("prints in defer foo func")
	}()
	fmt.Println("prints in defer foo")
}
//output is prints main,then defer foo and then defer foo func
