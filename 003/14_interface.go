package main

import (
	"fmt"
	"reflect"
)

type General interface{}

type GData interface {
	Set(nm string, g General) GData
	Print()
}

//NData
type NData struct {
	Name string
	Data []int
}

func (nd *NData) Set(nm string, g General) GData {
	nd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf(0)) { //型判定
		nd.Data = g.([]int) //型アサーション
	}
	return nd
}

func (nd *NData) Print() {
	fmt.Printf("<<%s>> value: %d\n", nd.Name, nd.Data)
}

//SData
type SData struct {
	Name string
	Data []string
}

func (sd *SData) Set(nm string, g General) GData {
	sd.Name = nm
	if reflect.TypeOf(g) == reflect.SliceOf(reflect.TypeOf("")) { //型の判定
		sd.Data = g.([]string) //型アサーション
	}
	return sd
}

func (sd *SData) Print() {
	fmt.Printf("* %s %s *\n", sd.Name, sd.Data)
}

func main() {
	data := []GData{}

	data = append(data, new(NData).Set("Taro", []int{123}))
	data = append(data, new(SData).Set("Jiro", []string{"hello"}))
	data = append(data, new(NData).Set("Hanako", []int{98700}))
	data = append(data, new(SData).Set("Sachiko", []string{"happy?"}))

	for _, ob := range data {
		ob.Print()
	}
}
