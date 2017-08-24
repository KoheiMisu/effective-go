package main

import (
	"time"
	"net/http"
	"fmt"
	"context"
	"log"
)


// ref http://go-talks.appspot.com/github.com/matope/talks/2016/context/context.slide

// contextパッケージ
// goの各種ライブラリへのキャンセル要求のインターフェイスの標準化

// できること
// goroutineやAPI教会をまたいだ処理のキャンセル要求と値の受け渡しが簡単に
func main() {
	req, _ := http.NewRequest("GET", "/health-check", nil)
	handle(req)
}

func slowProcess(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		log.Println("doing something...", i)
		select {
		case <-time.After(1*time.Second):
		case <-ctx.Done():
			log.Println("Slow process done", i)
			return ctx.Err()
		}
	}
	log.Println("Something is done")
	return nil
}

func handle(r *http.Request)  {
	// cancelを呼ぶとctxのDoneに送信が行われる

	//ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	//defer cancel()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-time.After(7 * time.Second)
		cancel()
	}()


	resultCh := make(chan error, 1)
	go func() {
		resultCh <- slowProcess(ctx)
	}()

	err := <-resultCh
	fmt.Println("Result", err)
	//fmt.Fprintln(w, "Result", err)
}