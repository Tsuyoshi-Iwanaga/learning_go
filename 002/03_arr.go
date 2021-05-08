package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := input("hello data")
	ar := strings.Split(x, " ")//半スペ区切りの数字を配列へ格納 10 20 30...
	t := 0

	// for i := 0; i < len(ar); i++ {
	// 	n, er := strconv.Atoi(ar[i])
	// 	if er != nil {
	// 		goto err
	// 	}
	// 	t += n
	// }

	for _, v := range ar { //rangeを使うと要素全てのループが回る(インデックスを管理しなくて良いので便利)
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}

	fmt.Println("total:", t)
	return

	err: fmt.Println("ERROR")
}

func input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}