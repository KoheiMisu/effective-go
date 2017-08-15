package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

// 終了時に何らかの処理を行って、結果を返す
// チャンネルで待ち受けて、処理の終了を待つようにする
func main() {
	rand.Seed(time.Now().UnixNano())

	quit := make(chan string)
	c := boring("Joe", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "Bye"
	fmt.Printf("Joe says %q\n", <-quit)
}

// goroutineを内部で動かす(受信専用チャンネルとして定義)
func boring(msg string, q chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s %d", msg, i):
			case <-q:
				cleanup()
				q <- "GoodBye"
				return
			}
		}
	}()

	return c // 受信チャンネルを返す
}

func cleanup() {

}
