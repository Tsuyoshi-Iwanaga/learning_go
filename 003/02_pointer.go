package main

import (
	"fmt"
)

func main() {
	n := 123
	m := 10000

	p := &n
	p2 := &m
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
}
