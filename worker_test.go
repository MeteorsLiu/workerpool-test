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

	for i := 0; i < 50; i++ {
		wg.Add(1)
		w.Schedule(func() {
			defer wg.Done()
			t.Log(i)
		})
	}
	wg.Wait()
}

func TestSeq(t *testing.T) {
	w := NewPool(1024, 1024, 1024)
	ch := make(chan int)
	for i := 0; i < 100; i++ {
		w.Schedule(func() {
			t.Log(<-ch)
		})
		ch <- i + 1
	}
	<-time.After(time.Minute)

}
