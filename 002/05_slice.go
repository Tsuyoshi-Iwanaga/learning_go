package main

import (
	"fmt"
)

func main() {
	a := [3]int{10, 20, 30}
	b := a[0:2]//aからスライスを生成

	fmt.Println(a)//[10, 20, 30]
	fmt.Println(b)//[10, 20]

	b = append(b, 1000)//スライスbの末尾に追加
	fmt.Println(a)//[10, 20, 1000]//元の配列に要素があれば書き換えられる
	fmt.Println(b)//[10, 20, 1000]

	b = append(b, 2000)
	fmt.Println(a)//[10, 20, 1000]//元の配列に要素がなければ追加はされない
	fmt.Println(b)//[10, 20, 1000, 2000]
}