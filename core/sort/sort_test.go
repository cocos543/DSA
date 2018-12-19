package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	// golang语法问题, 这里只能先把具体类型转成interface类型之后才能使用
	arr1 := []int{4, 5, 6, 3, 2, 1}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}

	BubbleSort(s, func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	})

	fmt.Println()

	arr2 := []string{"d", "e", "f", "c", "b", "a"}
	s = make([]interface{}, len(arr2))
	for i, v := range arr2 {
		s[i] = v
	}

	BubbleSort(s, func(a, b interface{}) bool {
		return a.(string) <= b.(string)
	})
}

func TestStraightInsertionSort(t *testing.T) {
	arr1 := []int{4, 5, 6, 1, 3, 2}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}

	StraightInsertionSort(s, func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	})
}
