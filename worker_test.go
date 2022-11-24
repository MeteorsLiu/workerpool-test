package worker

import (
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	w := NewPool(8, 1024, 8)
	for i := 0; i < 8; i++ {
		w.Schedule(func() {
			t.Log(i)
		})
	}
	<-time.After(time.Minute)
}
