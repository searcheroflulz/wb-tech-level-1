package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func worker(dataChannel chan int, wg *sync.WaitGroup, workerID int, ctx context.Context) {
	defer wg.Done()
	//fmt.Printf("Worker %d - starting session\n", workerID)

	for {
		select {
		case data := <-dataChannel:
			fmt.Printf("Worker %d - %v\n", workerID, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d - ending sessiom\n", workerID)
			return
		}
	}
}

func generateData(dataChannel chan int, ctx context.Context) {
	defer close(dataChannel)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			dataChannel <- rand.Intn(1000)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	dataChannel := make(chan int)

	var wg sync.WaitGroup
	numWorkers := 5

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(dataChannel, &wg, i, ctx)
	}
	go generateData(dataChannel, ctx)
	wg.Wait()

}
