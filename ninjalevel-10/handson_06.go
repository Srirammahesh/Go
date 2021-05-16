package main
import "fmt"
func main(){
	c:=make(chan int)

	go send(c)
	recieve(c)
	fmt.Println("exited")
}
func recieve(c <-chan int){
	for i :=range c {
		fmt.Println(i)
	}
}
func send(c chan<- int){
	for i:=0;i<100;i++ {
		c <- i
	}
	close(c)
}
