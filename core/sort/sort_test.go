package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1}
	BubbleSort(arr, func(a, b interface{}) bool {
		return a.(int) <= b.(int)
	})
}
