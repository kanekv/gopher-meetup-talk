package main

import (
	"bytes"
	"flag"
	"log"
	"runtime"
	"strings"
	"sync"
	"testing"
)

// SS OMIT
func NewBytesBufferPool() *sync.Pool {
	return &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}

// SE OMIT

// START OMIT
var bytz = []byte(strings.Repeat("*", 1000))
var bPool = NewBytesBufferPool()

func BenchmarkFreelist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		buf.Write(bytz)
		buf.Reset()
	}
}

func BenchmarkFreelist2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bPool.Get().(*bytes.Buffer)
		buf.Write(bytz)
		buf.Reset()
		bPool.Put(buf)
	}
}

// END OMIT

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{},
		[]testing.InternalBenchmark{{"BenchmarkFreelist", BenchmarkFreelist}, {"BenchmarkFreelist2", BenchmarkFreelist2}},
		[]testing.InternalExample{})
}
