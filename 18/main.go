package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/

type Counter struct {
	value int
	sync.RWMutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.value++
}

func main() {
	var wg sync.WaitGroup

	counter := Counter{}
	numGoroutines := 100

	incrementPerGoroutine := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println("\nCount is over:", counter.value)
}
