package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"//文字列⇄数値のキャストを行うモジュール
)

func main() {
	x := input("type a price")
	n, err := strconv.Atoi(x) //文字列(ASCII)を数値(Integer)へ変換
	// a := strconv.Itoa(n) 数値を文字列に変換

	if err != nil {
		fmt.Println("Error!")
		return
	}

	p := float64(n)//整数値を浮動小数点数へ変換
	fmt.Println(int(p * 1.1))//浮動小数点数を整数値に変換(切り捨て)
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}