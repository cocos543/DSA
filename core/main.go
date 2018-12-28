package main

import (
	. "DSA/core/sort"
	"fmt"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")

	arr1 := []int{11, 8, 3, 9, 7, 1, 2, 5}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	MergingSort(s, func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	})
}
