package main

import (
	"time"
	"log"
)

// channelを返して、そのチャンネルに送信することで
// 動かしてあるgoroutineが動くようになる

func worker(t time.Duration) chan int {
	// chan にバッファを設定
	ch := make(chan int, 20)
	go func() {
		for {
			time.Sleep(t * time.Millisecond)
			log.Println(<-ch)
		}
	}()
	return ch
}

// http://qiita.com/Jxck_/items/da3ca2db58734a966cac
func main() {
	workers := [](chan int){
		worker(10),
		worker(1000), // slow woker
		worker(10),
	}

	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		for _, w := range workers {
			select {
			case w <- i: // buffer が満杯だと実行されない
			default:
				// 空のデフォルトがあることでこの select を抜けられる。
				// もし無いと、 w が書き込み可能になるまで select がブロックする
			}
		}
	}

	// main を抜けないようにしてるだけ
	time.Sleep(10 * time.Second)
}

