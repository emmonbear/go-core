package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type CustomWaitGroup struct {
	ch    chan struct{}
	count int32
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{ch: make(chan struct{})}
}

func (cwg *CustomWaitGroup) Add(delta int32) {
	atomic.AddInt32(&cwg.count, delta)
	if cwg.count < 0 {
		panic("negative WaitGroup counter")
	}
	if atomic.LoadInt32(&cwg.count) == 0 {
		close(cwg.ch)
	}
}

func (cwg *CustomWaitGroup) Done() {
	cwg.Add(-1)
}

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
