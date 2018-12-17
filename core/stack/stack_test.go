package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s1 := NewStack(3)
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	fmt.Println(s1)
	assert.Equal(t, "1-2-3-", s1.String(), "Push is incorrect")
	assert.Equal(t, 3, s1.Count(), "Push is incorrect")
	assert.Equal(t, false, s1.Push(4), "Push is incorrect")

	assert.Equal(t, 3, s1.Pop(), "Push is incorrect")
	fmt.Println(s1)
	assert.Equal(t, 2, s1.Count(), "Push is incorrect")
	assert.Equal(t, "1-2-", s1.String(), "Push is incorrect")
	assert.Equal(t, 2, s1.Pop(), "Push is incorrect")
	assert.Equal(t, 1, s1.Pop(), "Push is incorrect")
	assert.Equal(t, nil, s1.Pop(), "Push is incorrect")
}

func TestLinkedStack(t *testing.T) {
	s1 := NewLikedStack(3)
	s1.Push(1)
	s1.Push(2)
	s1.Push(3)
	fmt.Println(s1)
	assert.Equal(t, "1-2-3-", s1.String(), "Push is incorrect")
	assert.Equal(t, 3, s1.Count(), "Push is incorrect")
	assert.Equal(t, false, s1.Push(4), "Push is incorrect")

	assert.Equal(t, 3, s1.Pop(), "Push is incorrect")
	fmt.Println(s1)
	assert.Equal(t, 2, s1.Count(), "Push is incorrect")
	assert.Equal(t, "1-2-", s1.String(), "Push is incorrect")
	assert.Equal(t, 2, s1.Pop(), "Push is incorrect")
	assert.Equal(t, 1, s1.Pop(), "Push is incorrect")
	assert.Equal(t, nil, s1.Pop(), "Push is incorrect")
}
