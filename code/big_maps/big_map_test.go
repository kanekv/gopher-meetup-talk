package main

import (
	"flag"
	"runtime"
	"testing"
)

// BEGIN OMIT
func TestLargeMapGC(t *testing.T) {
	m := make(map[string]*int, 10e6)
	i := 5
	m["foo"] = &i
	runtime.GC()
}

// END OMIT

func main() {
	flag.Set("test.v", "true")
	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestLargeMapGC", TestLargeMapGC}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
