package main

import (
	"flag"
	"log"
	"runtime"
	"testing"
)

// START OMIT
func TestMap(t *testing.T) {
	m := make(map[int]int, 1e6)
	for i := 0; i < 100; i++ {
		runtime.GC()
	}
	m[0] = 1
}

func TestSlice(t *testing.T) {
	m := make([]int, 1e6)
	for i := 0; i < 100; i++ {
		runtime.GC()
	}
	m[0] = 1
}

// END OMIT

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestMap", TestMap}, {"TestSlice", TestSlice}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
