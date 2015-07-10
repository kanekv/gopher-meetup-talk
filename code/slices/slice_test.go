package main

import (
	"flag"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// BEGIN OMIT
func TestSlice(t *testing.T) {
	slice := make([]string, 1)
	start := time.Now()
	for i := 0; i < 1e6; i++ {
		slice = append(slice, "1234567890abcdef")
	}
	println("TOTAL time string: ", time.Since(start)/1e6)
}

func TestSliceByte(t *testing.T) {
	sliceBytes := make([][16]byte, 1)
	start := time.Now()
	for i := 0; i < 1e6; i++ {
		sliceBytes = append(sliceBytes, [16]byte{})
	}
	println("TOTAL time bytes: ", time.Since(start)/1e6)
}

func BenchmarkSlice(b *testing.B) {
}

// END OMIT

func main() {
	fmt.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestSlice", TestSlice}, {"TestSliceByte", TestSliceByte}},
		[]testing.InternalBenchmark{{"BenchmarkSlice", BenchmarkSlice}},
		[]testing.InternalExample{})
}
