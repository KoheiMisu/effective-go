package main

import (
	"time"
	"math/rand"
	"fmt"
)

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%f result for %q\n", kind, query))
	}
}

//　処理の重さを考慮して80ms以下は処理を中断させる
func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) } ()
	go func() { c <- Image(query) } ()
	go func() { c <- Video(query) } ()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

// 低速サーバ上で破棄されないように処理の受け口を複製する
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	// searchReplicaで一番最初に受信したものをすぐに送信する
	// channelは送信時にもブロックを行う
	return <-c
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	// サーバーを複製 -> 複数のレプリカに要求を送信し、最初の応答を使用する
	results := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"),
	)
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}