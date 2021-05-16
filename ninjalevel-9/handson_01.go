package main
import (
	"fmt"
	"runtime"
	"sync"
)
var wg sync.WaitGroup

func main(){
	fmt.Println("number of CPUs:\t",runtime.NumCPU())
	fmt.Println("number of Goroutine:\t",runtime.NumGoroutine())
	wg.Add(2)
	go foo(4,4)
	go bar()
	wg.Wait()
	fmt.Println("number of Goroutine:\t",runtime.NumGoroutine())
}

func foo(a,b int){
	fmt.Println(a+b)
	wg.Done()

}

func bar(){
	for i:=0;i<10;i++{
		fmt.Println("bar",i)
	}
	wg.Done()
}