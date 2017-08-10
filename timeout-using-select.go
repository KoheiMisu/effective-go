package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 一つのgoroutineの処理時間がtimer以上だった場合に終了する
func main() {
	c := boring("Hey")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("You're too slow.")
			return
		}

	}
}

// goroutineを内部で動かす(受信専用チャンネルとして定義)
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			r := time.Duration(rand.Intn(2e3)) * time.Millisecond
			c <- fmt.Sprintf("%s %d time %d", msg, i, r)
			time.Sleep(r)
		}
	}()

	return c // 受信チャンネルを返す
}
