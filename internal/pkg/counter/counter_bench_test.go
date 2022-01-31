package counter

import (
	"sync"
	"testing"
)

func BenchmarkInc(b *testing.B) {
	c := NewCounter()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			Inc(c)
		}
	}
}

func BenchmarkIncPool(b *testing.B) {
	var bytesPool = sync.Pool{New: func() interface{} { return Counter{} }}
	bytesPool.Put(NewCounter())
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
			counter := bytesPool.Get().(*Counter)
			Inc(counter)
			bytesPool.Put(counter)
		}
	}
}
