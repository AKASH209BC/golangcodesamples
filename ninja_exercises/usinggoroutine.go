package main
import "fmt"
import "sync"
import "runtime"
var wg sync.WaitGroup

func main(){
	fmt.Println("GO OS Number CPU",runtime.NumCPU())
	fmt.Println("GO OS Architecture",runtime.NumGoroutine())
	fmt.Println("GO OS",runtime.GOOS)
	fmt.Println("GO Architecture",runtime.GOARCH)
	wg.Add(2)
	fmt.Println("Using go routine")
	go foobar()
	go foobar1()
	wg.Wait()
}
func foobar(){
	for i:=0;i<5;i++{
		fmt.Println("Hello from foobar",i)
	}
	wg.Done()
}

func foobar1(){
	for i:=0;i<5;i++{
		fmt.Println("Hello from foobar1",i)
	}
	wg.Done()
}
