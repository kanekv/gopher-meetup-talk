package main

import (
	"container/list"
	"flag"
	"log"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

// BEGIN OMIT
func TestLargeLinkedList(t *testing.T) {
	l := list.New()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10e6; i++ {
		l.PushBack(r.Int63())
	}
	runtime.GC()
}

func BenchmarkLinkedList(b *testing.B) {
	l := list.New()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < b.N; i++ {
		l.PushBack(r.Int63())
	}
}

// END OMIT

func main() {
	log.Println(runtime.Version())
	flag.Set("test.bench", "Large")
	flag.Set("test.v", "true")

	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{"TestLargeLinkedList", TestLargeLinkedList}},
		[]testing.InternalBenchmark{{"BenchmarkLinkedList", BenchmarkLinkedList}},
		[]testing.InternalExample{})
}
