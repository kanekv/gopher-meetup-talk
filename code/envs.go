package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPUs=", runtime.NumCPU())
	fmt.Println("NumCgoCall=", runtime.NumCgoCall())
	fmt.Println("NumGoroutines=", runtime.NumGoroutine())
	fmt.Println("Version=", runtime.Version())
	debug.PrintStack()
}
