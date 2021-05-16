package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)
//code contains race condition and curated using the atomic package
func main(){
	var wg sync.WaitGroup

	var ctr int64
	gs:=100
	wg.Add(gs)
	for i:=0;i<gs;i++ {
		go func() {
			atomic.AddInt64(&ctr,1)
			runtime.Gosched()
			fmt.Println("inside func",atomic.LoadInt64(&ctr))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
