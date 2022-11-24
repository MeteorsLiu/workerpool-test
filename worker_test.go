package worker

import (
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	w := NewPool(1, 1, 1)

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
