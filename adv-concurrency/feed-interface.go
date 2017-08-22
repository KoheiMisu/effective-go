package main

import "time"

type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

type Subscription interface {
	Updates() <-chan Item // stream of Items
	Close() error // shuts down the screen
}

type Item struct {
	Title, Channel, GUID string
}