package main

import (
	"log"
	"runtime"
	"runtime/debug"
)

func main() {
	log.Println("Version=", runtime.Version())
	log.Println("GOMAXPROCS=", runtime.GOMAXPROCS(0))
	log.Println("NumCPUs=", runtime.NumCPU())
	log.Println("NumCgoCall=", runtime.NumCgoCall())
	log.Println("NumGoroutines=", runtime.NumGoroutine())
	debug.PrintStack()
}
