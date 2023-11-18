package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

func sendData(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Intn(1000)
		}
		time.Sleep(5 * time.Second)
	}
}

func receiveData(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-ch:
			fmt.Printf("Receieved %d\n", data)
		}
	}
}

func main() {
	dataChan := make(chan int)
	sec := 30
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(sec)*time.Second)
	defer cancel()

	go sendData(ctx, dataChan)
	go receiveData(ctx, dataChan)
	<-ctx.Done()
	fmt.Println("Time is up")
}
