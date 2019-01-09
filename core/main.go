package main

import (
	"fmt"

	. "github.com/cocos543/DSA/core/sort"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")

	arr1 := []int{6, 11, 3, 9, 8}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	QuickSort(s, func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	})
}
