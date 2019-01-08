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

func TestBinarySearch1(t *testing.T) {
	arr1 := []int{0, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	assert.EqualValues(t, 1, BinarySearch1(s, 1, func(a, b interface{}) int { return a.(int) - b.(int) }), "BinarySearch1 is incorrect")
}

func TestBinarySearch2(t *testing.T) {
	arr1 := []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	assert.EqualValues(t, 3, BinarySearch2(s, 1, func(a, b interface{}) int { return a.(int) - b.(int) }), "BinarySearch2 is incorrect")
}

func TestBinarySearch3(t *testing.T) {
	arr1 := []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	assert.EqualValues(t, 1, BinarySearch3(s, 1, func(a, b interface{}) int { return a.(int) - b.(int) }), "BinarySearch3 is incorrect")
}

func TestBinarySearch4(t *testing.T) {
	arr1 := []int{0, 1, 1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := make([]interface{}, len(arr1))
	for i, v := range arr1 {
		s[i] = v
	}
	assert.EqualValues(t, 3, BinarySearch4(s, 1, func(a, b interface{}) int { return a.(int) - b.(int) }), "BinarySearch4 is incorrect")
}

func TestLetCodeBinarySearch(t *testing.T) {
	arr1 := []int{3, 1}
	assert.EqualValues(t, 1, LetCodeBinarySearch(arr1, 1), "LetCodeBinarySearch is incorrect")
}
