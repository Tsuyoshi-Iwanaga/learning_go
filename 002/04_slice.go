package main

import (
	"fmt"
)

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]//aからスライスを生成

	fmt.Println(a)//[10, 20, 30, 40, 50]
	fmt.Println(b)//[10, 20, 30]

	a[0] = 100
	fmt.Println(a)//[100, 20, 30, 40, 50]
	fmt.Println(b)//[100, 20, 30]元配列の変更が影響する

	b[1] = 200
	fmt.Println(a)//[100, 200, 30, 40, 50]スライスの変更が影響する
	fmt.Println(b)//[100, 200, 30]
}