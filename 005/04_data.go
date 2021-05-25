package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rt := func(f *os.File) {
		//Readerを作成する
		r := bufio.NewReaderSize(f, 4096)
		for i := 1; true; i++ {
			s, _, er := r.ReadLine()
			if er != nil {
				break
			}
			fmt.Println(i, ":", string(s))
		}
	}

	f, er := os.OpenFile("02_data.txt", os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	defer f.Close()

	fmt.Println("<< start >>")
	rt(f)
	fmt.Println("<< end >>")
}
