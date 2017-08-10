package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You've both boring; i'm leaving")
}

// goroutineを内部で動かす(受信専用チャンネルとして定義)
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			r := rand.Intn(500) + 100
			time.Sleep(time.Duration(r) * time.Millisecond)
			c <- fmt.Sprintf("%s %d time %d", msg, i, r)
		}
	}()

	return c // 受信チャンネルを返す
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
