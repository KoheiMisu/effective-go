package main

import (
	"time"
	"net/http"
	"log"
	"context"
	"fmt"
	"io/ioutil"
)

// Requestにはcontextが実装されたので終了指示がcontextで上書きできるようになった
func main()  {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	//ctx, cancel := context.WithTimeout(req.Context(), 100*time.Millisecond)
	ctx, cancel := context.WithTimeout(req.Context(), 1*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if b, err := ioutil.ReadAll(res.Body); err == nil {
		fmt.Println(string(b))
	}

}