package main

import "fmt"

func total(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t //チャンネルに値を渡す
}

func main() {
	c := make(chan int) //チャンネルを作成
	go total(1000, c)
	go total(100, c)
	go total(10, c)
	x, y, z := <-c, <-c, <-c //チャンネルから値を取り出す(先入れ先出す)
	fmt.Println(x, y, z)
}
