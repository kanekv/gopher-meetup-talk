package main

import (
	"flag"
	"fmt"
	"runtime"
	"testing"
)

// START OMIT
type mapStruct struct {
	userId int
	b      [16]byte
}

func TestMap(t *testing.T) {
	m := make(map[int]*mapStruct, 1e6)
	for i := 0; i < 20e6; i++ {
		m[i] = &mapStruct{}
	}
}

func TestMap2(t *testing.T) {
	m := make(map[int]mapStruct, 1e6)
	for i := 0; i < 20e6; i++ {
		m[i] = mapStruct{}
	}
}

// END OMIT

func main() {
	fmt.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestMap", TestMap}, {"TestMap2", TestMap2}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
