package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// слушаем канал, по которому может что-то прийти, и если приходит - выходим из горутины.
func channelTermination(stopChannel chan bool) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Received a signal")
			return
		default:
			fmt.Println("Everything is fine")
			time.Sleep(time.Second)
		}
	}
}

// передаем в горутину указатель на общую переменную. Если она меняется - выходим из горутины.
func sharedVariable(stopIt *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if *stopIt {
			fmt.Println("Received a shared true")
			return
		}

		fmt.Println("Everything is fine")
		time.Sleep(time.Second)
	}
}

// получаем закрытие канала - выходим из горутины.
func channelClosure(stopChannel chan struct{}) {
	for {
		select {
		case <-stopChannel:
			fmt.Println("Received a closure")
			return
		default:
			fmt.Println("Everything is fine")
			time.Sleep(time.Second)
		}
	}
}

// контекст отменяется по какой-либо причине - выходим из горутины.
func contextNotification(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Received cancel")
			return
		default:
			fmt.Println("Everything is fine")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	{
		stopChannel := make(chan bool)
		go channelTermination(stopChannel)
		time.Sleep(3 * time.Second)
		stopChannel <- true
		time.Sleep(time.Second)
	}
	time.Sleep(30 * time.Second)
	{
		var stopIt bool
		var wg sync.WaitGroup
		wg.Add(1)
		go sharedVariable(&stopIt, &wg)
		time.Sleep(3 * time.Second)
		stopIt = true
		wg.Wait()
	}
	time.Sleep(30 * time.Second)
	{
		stopChannel := make(chan struct{})
		go channelClosure(stopChannel)
		time.Sleep(3 * time.Second)
		close(stopChannel)
		time.Sleep(time.Second)
	}
	time.Sleep(30 * time.Second)
	{
		ctx, cancel := context.WithCancel(context.Background())
		go contextNotification(ctx)
		time.Sleep(3 * time.Second)
		cancel()
		time.Sleep(time.Second)
	}
}
