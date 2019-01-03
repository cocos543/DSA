package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	arr1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	assert.EqualValues(t, 0, BinarySearch(s, 0, func(a, b interface{}) int { return a.(int) - b.(int) }), "BinarySearch is incorrect")
}

func TestSquare(t *testing.T) {
	fmt.Println(Square(3))
}
