package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n=", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total:", t)
}

func main() {
	c := make(chan int) //チャンネル作成
	go total(c)         //一旦チャンネル渡してGoルーチンを実行する
	c <- 100            //後からチャンネルに値を流す(Goルーチン実行より前だとエラーになる)
	time.Sleep(100 * time.Millisecond)
}
