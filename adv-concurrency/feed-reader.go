package main

import (
	"time"
	"fmt"
)

// sub implements the Subscription interface
type sub struct {
	naiveSub

	fetcher Fetcher
	updates chan Item
}

type naiveSub struct {
	closed bool
	err error
}

// loop fetches items using s.fetcher and sends them on s.updates
// loop exists when s.Close is called

// fetchを定期的に呼び出す
// updateチャンネルで取得したアイテムを送信する
// Closeが呼び出されたときに終了し、エラーを報告する

// select-forで管理
// Closeが呼ばれた時
// fetchできるようになったとき
// s.updatesにitemを送る
func (s *sub) loop() {


	for {
		if s.closed {
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err
			time.Sleep(10 * time.Second)
			continue
		}
		for _, item := range items {
			s.updates <- item
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now))
		}
	}
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	// Todo: make loop exit
	// Todo: find out about any error
	return err
}

func (s *naiveSub) Close() error {
	s.closed = true
	return s.err
}


// converts Fetches to a stream
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	go s.loop()
	return s
}

// merges several streams
func Merge(subs ...Subscription) Subscription {

}

// fetch Items from domain
func Fetch(domain string) Fetcher {

}

func main() {
	// subscribe to some feeds, and create a merged update stream
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")),
	)

	// Close the subscriptions after some time
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed: ", merged.Close())
	})

	// Print the screen
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	panic("show me the stacks")
}