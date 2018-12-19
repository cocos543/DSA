package main

import (
	. "DSA/core/sort"
	"fmt"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")

	arr := []int{1, 2, 3}
	Swap(arr, 1, 2)
	fmt.Println(arr)
}
