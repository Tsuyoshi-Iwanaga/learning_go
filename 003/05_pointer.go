package main

import "fmt"

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initial(&ar)
	fmt.Println(ar)
}

func initial(ar *[]int) {
	for i := range *ar {
		(*ar)[i] = 0
	}
}
