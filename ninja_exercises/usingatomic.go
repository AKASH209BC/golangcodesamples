package main

import "fmt"
import "sync"
import "sync/atomic"

func main(){
	var wg sync.WaitGroup
	var c int64
	c1:= 100
	wg.Add(c1)
	for i:=0;i<c1;i++ {
		go func() {
			atomic.AddInt64(&c, 1)
			fmt.Println(atomic.LoadInt64(&c))
			wg.Done()
		}()
	}
	wg.Wait()
}
