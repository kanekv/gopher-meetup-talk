package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

// BEGIN OMIT
func TestPipe(t *testing.T) {
	var b = []byte(strings.Repeat("*", 1000))
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		r, w := io.Pipe()
		go func() {
			for i := 0; i < 100e3; i++ {
				w.Write(b)
			}
			w.Close()
		}()

		buf := &bytes.Buffer{}
		io.Copy(buf, r)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("TOTAL time Pipe: ", time.Since(start)/1e6)
}

// END OMIT

func main() {
	fmt.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestPipe", TestPipe}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
