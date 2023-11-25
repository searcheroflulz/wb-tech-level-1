package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func detectType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	default:
		fmt.Printf("unknown type: %v", reflect.TypeOf(variable))
	}
}

func main() {
	var integer int
	detectType(integer)

	var string string
	detectType(string)

	var bool bool
	detectType(bool)

	channel := make(chan int)
	detectType(channel)

	arr := make([]int, 10)
	detectType(arr)
}
