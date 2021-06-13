package main

import "fmt"
import "runtime"
import "sync"

var wg sync.WaitGroup

func main(){
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	counter:=0
	wg.Add(2)
	go porter1(counter)
	go porter2(counter)
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}

func porter1(c int){
	var m sync.Mutex
	m.Lock()
	v:= c
	v++
	c=v
	m.Unlock()
	wg.Done()
}

func porter2(c int){
	var m sync.Mutex
	m.Lock()
	v:= c
	v++
	c=v
	m.Unlock()
	wg.Done()
}
