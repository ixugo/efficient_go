package gc

import (
	"runtime"
	"testing"
	"time"
)

func BenchmarkTimer(b *testing.B) {
	b.Run("timer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				t := time.NewTimer(5 * time.Microsecond)
				select {
				case <-t.C:
				default:
					t.Stop()
				}

			}()
		}
	})

	b.Run("after", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				select {
				case <-time.After(5 * time.Microsecond):
				default:
				}
			}()
		}
	})
}

func BenchmarkTimer2(b *testing.B) {
	runtime.GC()
	b.Run("timer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				t := time.NewTimer(5 * time.Microsecond)
				t.Stop()
				t = nil

			}()
		}
	})

	b.Run("timer2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			func() {
				t := time.NewTimer(5 * time.Microsecond)
				t.Stop()
			}()
		}
	})
}
