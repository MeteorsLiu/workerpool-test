package worker

import (
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

	ch := make(chan int)
	for i := 0; i < 50; i++ {
		w.Schedule(func() {
			t.Log(<-ch)
		})
		ch <- i + 1
	}
	<-time.After(time.Minute)
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
