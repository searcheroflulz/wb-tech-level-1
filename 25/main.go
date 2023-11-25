package main

import (
	"fmt"
	"time"
)

/*
Реализовать собственную функцию sleep.
*/

// ожидаем закрытия канала, который закроется после указанного количества секунд
func sleep(duration time.Duration) {
	<-time.After(duration * time.Second)
}

func main() {
	fmt.Println(time.Now())
	sleep(5)
	fmt.Println(time.Now())
}
