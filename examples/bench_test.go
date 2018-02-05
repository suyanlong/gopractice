package main

import (
	"testing"
	"sync"
	"sync/atomic"
)

func BenchmarkTest_SyncMap(b *testing.B) {
	kk := sync.Map{}
	go func() {
		for i := 0; i < b.N; i++ {
			kk.Store(i, i)
		}
	}()

	for i := 0; i < b.N; i++ {
		kk.Store(i, i)
	}

}

func BenchmarkTest_MutexMap(b *testing.B) {
	kk := make(map[int]int)
	mu := sync.Mutex{}
	go func() {
		for i := 0; i < b.N; i++ {
			mu.Lock()
			kk[i] = i
			mu.Unlock()
		}
	}()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		kk[i] = i
		mu.Unlock()
	}

}

func BenchmarkTest_atomic_int64(b *testing.B) {
	var kk int64 = 0
	go func() {
		for i := 0; i < b.N; i++ {
			atomic.StoreInt64(&kk, int64(i))
		}

	}()

	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&kk, int64(i))
	}

}

func BenchmarkTest_mutex_int64(b *testing.B) {
	var kk int64 = 0
	mu := sync.Mutex{}
	go func() {
		for i := 0; i < b.N; i++ {
			mu.Lock()
			kk = int64(i)
			mu.Unlock()
		}
	}()

	for i := 0; i < b.N; i++ {
		mu.Lock()
		kk = int64(i)
		mu.Unlock()
	}
}

func BenchmarkTest_atomic(b *testing.B) {
	var kk int64 = 0
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&kk, int64(i))
	}
}

func BenchmarkTest_int64(b *testing.B) {
	var kk int64 = 0
	for i := 0; i < b.N; i++ {
		kk += int64(i)
	}
}
