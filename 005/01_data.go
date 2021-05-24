package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//read File
	ba, er := ioutil.ReadFile("./01_data.txt")
	if er == nil {
		fmt.Println(string(ba))
	} else {
		fmt.Println("Read Error")
	}

	fmt.Println("")

	//write File(OverWrite)
	err := ioutil.WriteFile("./01_data2.txt", []byte("hello go!"), os.ModePerm)
	if err == nil {
		fmt.Println("Write finish")
	} else {
		fmt.Println("Write Error")
	}
}
