package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			st := time.Duration(rand.Intn(500)+100) * time.Millisecond
			time.Sleep(st)
			fmt.Println(st)
			wg.Done()
		}()
	}
	wg.Wait()
}
