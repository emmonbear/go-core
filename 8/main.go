package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// CustomWaitGroup is a custom implementation of a WaitGroup with functionality
// similar to the standard sync.WaitGroup but using channels and atomic operations
// for managing the goroutine count.
type CustomWaitGroup struct {
	ch    chan struct{}
	count int32
}

// NewCustomWaitGroup creates a new instance of CustomWaitGroup.
func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{ch: make(chan struct{})}
}

// Add increments or decrements the count by the given delta.
// It closes the internal channel when the count reaches zero.
func (cwg *CustomWaitGroup) Add(delta int32) {
	atomic.AddInt32(&cwg.count, delta)
	if cwg.count < 0 {
		panic("negative WaitGroup counter")
	}
	if atomic.LoadInt32(&cwg.count) == 0 {
		close(cwg.ch)
	}
}

// Done decrements the count by 1, typically called when a goroutine finishes.
func (cwg *CustomWaitGroup) Done() {
	cwg.Add(-1)
}

// Wait blocks until the count reaches zero, indicating all goroutines are done.
func (cwg *CustomWaitGroup) Wait() {
	<-cwg.ch
}

func main() {
	cwg := NewCustomWaitGroup()
	for i := 0; i < 10; i++ {
		cwg.Add(1)
		go func(index int) {
			defer cwg.Done()
			fmt.Printf("worker %d starting\n", i)
			time.Sleep(time.Second)
			fmt.Printf("worker %d done\n", i)
		}(i)
	}
	cwg.Wait()
	fmt.Println("All workers done")
}
