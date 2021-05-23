package main

import "fmt"

func main() {
	n := 123
	fmt.Printf("value: %d.\n", n)
	change1(n)
	fmt.Printf("value: %d.\n", n)
	change2(&n)
	fmt.Printf("value: %d.\n", n)
}

//仮引数には実引数の「コピー」が渡される(値渡し)
func change1(n int) {
	n *= 2
}

//仮引数には実引数の「ポインタ」が渡される(参照渡し、ポインタ渡し)
//これにより呼び出し元の値を変更できる
func change2(n *int) {
	*n *= 2
}
