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
		fmt.Println(kind)
		return Result(fmt.Sprintf("%f result for %q\n", kind, query))
	}
}

// 処理の重さを考慮して80ms以下は処理を中断させる
// 低速サーバの場合
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

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}