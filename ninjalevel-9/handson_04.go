package main

import (
	"fmt"
	"runtime"
	"sync"
)

//code contains race condition and curated using mutex
func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	ctr:=0
	gs:=100
	wg.Add(gs)
	for i:=0;i<gs;i++ {
		go func() {
			mu.Lock()
			v := ctr

			runtime.Gosched()
			v++
			ctr = v
			fmt.Println("inside func",ctr)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
