package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(1)
	queue.EnQueue(1)
	fmt.Println(queue)
	assert.Equal(t, "1-", queue.String(), "EnQueue is incorrect")
	assert.Equal(t, false, queue.EnQueue(2), "EnQueue is incorrect")

	queue = NewQueue(3)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	fmt.Println(queue)
	assert.Equal(t, "1-2-3-", queue.String(), "EnQueue is incorrect")
	assert.Equal(t, 1, queue.DeQueue(), "DeQueue is incorrect")
	fmt.Println(queue)
	assert.Equal(t, "2-3-", queue.String(), "EnQueue is incorrect")
	queue.DeQueue()
	queue.DeQueue()
	fmt.Println(queue)
	assert.Equal(t, "-", queue.String(), "EnQueue is incorrect")
	assert.Equal(t, nil, queue.DeQueue(), "DeQueue is incorrect")
}
