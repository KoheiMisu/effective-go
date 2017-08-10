package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go boring("boring", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You've boring; I'm leaving.")
}

// 都度メッセージを送信する
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		time.Sleep(1000 * time.Millisecond)
		c <- fmt.Sprintf("%s %d", msg, i)
	}
}
