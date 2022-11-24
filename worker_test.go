package worker

import (
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	w := NewPool(1, 1024, 1)

	w.Schedule(func() {
		t.Log(1)
	})
	w.Schedule(func() {
		t.Log(2)
	})
	w.Schedule(func() {
		t.Log(3)
	})
	w.Schedule(func() {
		t.Log(4)
	})
	w.Schedule(func() {
		t.Log(5)
	})
	w.Schedule(func() {
		t.Log(6)
	})
	w.Schedule(func() {
		t.Log(7)
	})
	w.Schedule(func() {
		t.Log(8)
	})

	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(50)
	for i := 0; i < 50; i++ {
		w.Schedule(func() {
			defer wg.Done()
			t.Log(<-ch)
		})
		ch <- i + 1
	}
	wg.Wait()
}

func TestSeq(t *testing.T) {
	w := NewPool(1024, 1024, 1024)
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		w.Schedule(func() {
			defer wg.Done()
			t.Log(<-ch)
		})
		ch <- i + 1
	}
	wg.Wait()

}

func BenchmarkNaive(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
		}()
	}
	wg.Wait()
}

func BenchmarkPool(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(b.N)
	w := NewPool(10000, 10000, 10000)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Schedule(func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
		})
	}
	wg.Wait()
}
