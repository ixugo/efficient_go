package maps

import (
	"strconv"
	"sync"
	"testing"

	"github.com/easierway/concurrent_map"
)

type Map interface {
	Set(k, v any)
	Get(any) (any, bool)
	Del(any)
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup

		// 并发 100 个协程写入
		for i := 0; i < WRITER; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100; i++ {
					hm.Set(strconv.Itoa(i), i*i)
					hm.Set(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
			}()
		}
		// 并发 100 个协程查询
		for i := 0; i < READER; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
			}()
		}
		wg.Wait()
	}
}

const (
	READER = 1000
	WRITER = 10
)

func BenchmarkSyncmap(b *testing.B) {
	// 性能较差
	b.Run("RWLock map", func(b *testing.B) {
		hm := createRWLockMap(1000)
		benchmarkMap(b, hm)
	})
	// 适合读多写少
	b.Run("sync.map", func(b *testing.B) {
		m := concurrent_map.CreateSyncMapBenchmarkAdapter()
		benchmarkMap(b, m)
	})
	// 适合写多读少
	b.Run("concurrent map", func(b *testing.B) {
		m := concurrent_map.CreateConcurrentMapBenchmarkAdapter(1000)
		benchmarkMap(b, m)
	})

}
