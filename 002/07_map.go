package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}
	m["total"] = m["a"] + m["b"] + m["c"]
	fmt.Println(m)

	//mapから特定のキーを削除
	delete(m, "total")

	//mapをループで走査する
	for k, v := range m {
		fmt.Println(k + ":", v)
	}
}