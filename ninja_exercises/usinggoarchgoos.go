package main

import (
"fmt"
"runtime"
)

func main() {
	// use go build and go install to compile it
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
