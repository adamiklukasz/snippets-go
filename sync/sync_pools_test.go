package sync

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Printf("New buffer created\n")
		return &bytes.Buffer{}
	},
}

func writeWithPool(iMax, jMax int) {
	a := "abcde"
	for i := 0; i <= iMax; i++ {

		bf := bufferPool.Get().(*bytes.Buffer)
		bf.Reset()

		for j := 0; j <= jMax; j++ {
			bf.WriteString(a)
		}

		bufferPool.Put(bf)
	}
}

func writeWithoutPool(iMax, jMax int) {
	a := "abcde"
	for i := 0; i <= iMax; i++ {

		bf := &bytes.Buffer{}

		for j := 0; j <= jMax; j++ {
			bf.WriteString(a)
		}
	}
}

//BenchmarkWithPool_10_1000-4                18998             67541 ns/op
//BenchmarkWithoutPool_10_1000-4              8175            203877 ns/op

func BenchmarkWithPool_10_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeWithPool(10, 1000)
	}
}

func BenchmarkWithoutPool_10_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		writeWithoutPool(10, 1000)
	}
}
