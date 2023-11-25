package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

// создааем структуру Human
type Human struct {
	Age     int
	Name    string
	Address string
}

func (h *Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

// в структуре Action реализуем встраивание
type Action struct {
	Human
}

func (a *Action) Speak() {
	fmt.Println("I'm", a.Age, "old")
}

func main() {
	person := Human{
		Age:     18,
		Name:    "John",
		Address: "Johnson street",
	}
	action := Action{person}
	action.Walk()
	action.Speak()
}
