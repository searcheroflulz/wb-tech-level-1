package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Реализовать конкурентную запись данных в map.
*/

func writeData(dataMap map[int]int, mutex *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	dataMap[rand.Intn(100)]++
	mutex.Unlock()
}

func main() {
	dataMap := make(map[int]int)
	var mutex sync.RWMutex

	var wg sync.WaitGroup
	numGoroutines := 9

	for i := 0; i <= numGoroutines; i++ {
		wg.Add(1)
		go writeData(dataMap, &mutex, &wg)
	}
	wg.Wait()
	fmt.Println(dataMap)
}
