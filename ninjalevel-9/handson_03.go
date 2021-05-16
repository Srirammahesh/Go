package main

import (
	"fmt"
	"runtime"
	"sync"
)
//code contains race condition
func main(){
	var wg sync.WaitGroup

	ctr:=0
	gs:=100
	wg.Add(gs)
	for i:=0;i<gs;i++ {
		go func() {
			v := ctr

			runtime.Gosched()
			v++
			ctr = v
			fmt.Println("inside func",ctr)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
