package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rt := func(f *os.File) {
		s, er := ioutil.ReadAll(f)
		if er != nil {
			panic(er)
		}
		fmt.Println(string(s))
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
