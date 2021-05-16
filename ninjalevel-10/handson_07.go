package main

import (
	"fmt"
)
func main(){
	//var wg sync.WaitGroup
	const goroutines=10
	//wg.Add(goroutines)
	c:=make(chan int)
	for i:=0;i<goroutines;i++ {
		go func() {
			for i:=0;i<10;i++{
				c<-i
			}
		}()
		//wg.Done()

	}
	for k:=0;k<100;k++{
		fmt.Println(<-c)
	}
	//wg.Wait()
	fmt.Println("exited")
}
