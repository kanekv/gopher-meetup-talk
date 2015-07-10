package main

import (
	"flag"
	"log"
	"runtime"
	"runtime/debug"
	"testing"
	"time"
)

// BEGIN OMIT
func TestSlice(t *testing.T) {
	debug.SetGCPercent(-1)
	slice := make([]string, 1e6)
	start := time.Now()
	for i := 0; i < 10e6; i++ {
		slice = append(slice, "1234567890abcdef")
	}
	log.Println("TOTAL time string: ", time.Since(start)/1e6)
}

func TestSliceByte(t *testing.T) {
	sliceBytes := make([][16]byte, 1e6)
	start := time.Now()
	for i := 0; i < 10e6; i++ {
		sliceBytes = append(sliceBytes, [16]byte{})
	}
	log.Println("TOTAL time bytes: ", time.Since(start)/1e6)
}

// END OMIT

func BenchmarkSlice(b *testing.B) {
}

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestSlice", TestSlice}, {"TestSliceByte", TestSliceByte}},
		[]testing.InternalBenchmark{{"BenchmarkSlice", BenchmarkSlice}},
		[]testing.InternalExample{})
}
