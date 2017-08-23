package main

import (
	"time"
	"net/http"
	"fmt"
)


// ref http://go-talks.appspot.com/github.com/matope/talks/2016/context/context.slide

// contextパッケージ
// goの各種ライブラリへのキャンセル要求のインターフェイスの標準化

// できること
// goroutineやAPI教会をまたいだ処理のキャンセル要求と値の受け渡しが簡単に
func main() {
	w := http
	handle()
}

func slowProcess() error {
	for i := 0; i < 10; i++ {
		fmt.Println("doing something ...")
		time.Sleep(time.Duration(1 * time.Second))
	}
	fmt.Println("something is done")
	return nil
}

func handle(w http.ResponseWriter, r *http.Request)  {
	resultCh := make(chan error, 1)
	go func() {
		resultCh <- slowProcess()
	}()

	select {
	case err := <-resultCh:
		fmt.Fprintln(w, "Result", err)
	}
}