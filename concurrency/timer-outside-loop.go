package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// timer設定時刻を超えたら終了する場合
func main() {
	c := boring("Hey")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
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
