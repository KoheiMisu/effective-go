package main

import (
	"fmt"
	"sync"
)

// メモリを捨てずに再利用することができる
// メモリアロケーション(割り当て)の最適化を図れる
var pool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 10000)
	},
}

func main() {
	// 在庫がなければNewが呼ばれて新しいものが作られる
	c := pool.Get().([]int)
	fmt.Println(c, len(c), cap(c))

	c = append(c, 1, 2, 3, 4)
	pool.Put(c)

	c = pool.Get().([]int)
	fmt.Println(c, len(c), cap(c)) // 値を使いまわしているので[1 2 3 4]

	// Todo 上記のような理由から、Resetメソッドを作っておいたほうが安心して使えそう
}
