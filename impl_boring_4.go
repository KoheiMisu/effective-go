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

// Todo 可変長引数で動くようにしたい
// 受信側は任意の数の実行手順を受け取って実行するようにする
// というか、送信側と受信側は1:1でその対を実行する分増やす方がよい??
func fanIn(input ...<-chan string) <-chan string {
	c := make(chan string)

	for _, i := range input {
		go func() {
			for {
				c <- <-i
			}
		}()
	}
	return c
}

//上記だと全てAnnで直列実行になってしまう

//Ann 0 time 487
//Ann 1 time 447
//Ann 2 time 159
//Ann 3 time 181
//Ann 4 time 418
//Ann 5 time 525
//Ann 6 time 140
//Ann 7 time 556
//Ann 8 time 400
//Ann 9 time 294
//You've both boring; i'm leaving
