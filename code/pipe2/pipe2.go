package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

// SETUP_START OMIT
var b = []byte(strings.Repeat("*", 1000))

type CustomPipe struct {
	buf bytes.Buffer
}

func (br *CustomPipe) Read(p []byte) (n int, err error) {
	return 0, nil
}
func (br *CustomPipe) readmore() (n int, err error) {
	return br.buf.Write(b)
}
func (br *CustomPipe) WriteTo(dst io.Writer) (n int64, err error) {
	n = 0
	for i := 0; i < 100e3; i++ {
		br.readmore()
		written, err := br.buf.WriteTo(dst)
		if err != nil {
			return 0, err
		}
		n += written
	}
	return n, err
}

// SETUP_END OMIT

// BEGIN OMIT
func TestPipe2(t *testing.T) {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		br := &CustomPipe{}
		buf := &bytes.Buffer{}
		io.Copy(buf, br)
		wg.Done()
	}()
	wg.Wait()
	log.Println("TOTAL time Pipe2: ", time.Since(start)/1e6)
}

// END OMIT

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "foo")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestPipe2", TestPipe2}},
		[]testing.InternalBenchmark{},
		[]testing.InternalExample{})
}
