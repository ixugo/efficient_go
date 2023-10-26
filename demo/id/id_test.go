package id

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/sony/sonyflake"
)

func TestID(t *testing.T) {
	start := time.Unix(1672502400, 0)
	cache := make(map[uint64]struct{})
	s := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: start,
		MachineID: func() (uint16, error) {
			return 1, nil
		},
	})
	s2 := sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime: start,
		MachineID: func() (uint16, error) {
			return 2, nil
		},
	})
	ch := make(chan uint64, 10)
	go func() {
		for v := range ch {
			_, ok := cache[v]
			if ok {
				panic("exist")
			}
			cache[v] = struct{}{}
			fmt.Println(v)
		}

	}()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			id, err := s.NextID()
			if err != nil {
				panic(err)
			}
			ch <- id

		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			id, err := s2.NextID()
			if err != nil {
				panic(err)
			}
			ch <- id
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			id, err := s.NextID()
			if err != nil {
				panic(err)
			}
			ch <- id
		}
	}()
	wg.Wait()

}
