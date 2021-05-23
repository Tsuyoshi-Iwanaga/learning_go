package main

import "fmt"

//空のインターフェース、どんな値でも入れることができる
type General interface{}

type GData interface {
	Set(nm string, g General)
	Print()
}

type GDataImpl struct {
	Name string
	Data General
}

func (gd *GDataImpl) Set(nm string, g General) {
	gd.Name = nm
	gd.Data = g
}

func (gd *GDataImpl) Print() {
	fmt.Printf("<<%s>> ", gd.Name)
	fmt.Println(gd.Data)
}

func main() {
	var data = []GDataImpl{}
	data = append(data, GDataImpl{"Taro", 123})
	data = append(data, GDataImpl{"Hanako", "hello"})
	data = append(data, GDataImpl{"Sachiko", []int{123, 456, 789}})
	for _, ob := range data {
		ob.Print()
	}
}
