package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	joe := boring("Joe")
	ann := boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
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
