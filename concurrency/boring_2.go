package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("boring!") // 関数がチャンネルを返すようにする
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving")
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
