package main

import (
	. "DSA/core/sort"
	"fmt"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")

	arr2 := []string{"a", "c", "b", "f", "e"}
	s := make([]interface{}, len(arr2))
	for i, v := range arr2 {
		s[i] = v
	}

	BubbleSort(s, func(a, b interface{}) bool {
		return a.(string) <= b.(string)
	})
}
