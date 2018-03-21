package goroutine

import (
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	goroutine()
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			time.Sleep(time.Microsecond * 1000)
		}()
		
	}
}
