package main

import (
	"math/rand"
	"fmt"
)

func main() {
	a, b := make(chan string), make(chan string)

	go func() { a <- "a" }()
	go func() { b <- "b" }()

	if rand.Intn(2) == 0 {
		a = nil
		fmt.Println("nil a")
	} else {
		b = nil
		fmt.Println("nil b")
	}
	select {
	case s := <-a:
		fmt.Println("got", s)
	case s := <-b:
		fmt.Println("got", s)
	}
}