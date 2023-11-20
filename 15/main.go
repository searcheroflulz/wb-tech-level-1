package main

import (
	"strings"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

*/

var justString string

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	/*justString = v[:100]
	данный фрагмент кода может привести к тому, что произойдет утечка памяти из-за того,
	что оригинальная строка (v) так и будет лежать в памяти из-за ссылки на нее в переменной justString.
	поэтому лучше сделать отдельный слайс байтов, в который мы скопируем лишь нужную часть из переменной v.
	таким образом глобальная переменная justString будет ссылаться только на необходимые нам данные, а переменная v освободится из памяти
	*/
	newSlice := make([]byte, 100)
	copy(newSlice, v[:100])
	justString = string(newSlice)
}

func main() {
	someFunc()
}
