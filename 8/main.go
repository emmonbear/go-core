package main

import (
	"fmt"
	"sync"
	"time"
)

type CustomWaitGroup struct {
	semaphore chan struct{}
	counter   int
	mutex sync.Mutex
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		semaphore: make(chan struct{}),
	}
}

func (wg *CustomWaitGroup) Add(delta int) {
	wg.mutex.Lock()
	defer wg.mutex.Unlock()
	wg.counter += delta
	if wg.counter == 0 {
		close(wg.semaphore)
	}
}

func (wg *CustomWaitGroup) Done() {
	wg.Add(-1)
}

func (wg *CustomWaitGroup) Wait() {
	<-wg.semaphore
}

func main() {
	wg := NewCustomWaitGroup()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Worker %d starting\n", id)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("All workers done")
}
