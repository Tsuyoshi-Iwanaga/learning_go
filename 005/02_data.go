package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//書き込みを行う関数
	wt := func(f *os.File, s string) {
		_, er := f.WriteString(s + "\n")
		if er != nil {
			panic(er)
		}
	}

	//ファイルポインタを取得
	f, er := os.OpenFile("02_data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	//defer close.
	defer f.Close()

	fmt.Println("*** start ***")
	wt(f, "*** start ***")

	for {
		s := input("type message")
		if s == "" {
			break
		}
		wt(f, s)
	}

	wt(f, "*** end ***\n\n")
	fmt.Println("*** end ***")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
