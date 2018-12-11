package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertList(t *testing.T) {
	list := NewSinglyLinkedList(nil)
	for i := 0; i < 10; i++ {
		node := NewSinglyLinkedNode(i + 1)
		list.InsertNode(node)
	}
	list.String()
	assert.Equal(t, 1, list.head.next.GetValue(), "InsertNode is incorrect")

	list = NewSinglyLinkedList(nil)
	for i := 0; i < 10; i++ {
		node := NewSinglyLinkedNode(i + 1)
		list.InsertNodeHead(node)
	}
	list.String()

	list = NewSinglyLinkedList(nil)
	node1 := NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	node3 := NewSinglyLinkedNode(3)
	node4 := NewSinglyLinkedNode(4)
	node5 := NewSinglyLinkedNode(5)
	node6 := NewSinglyLinkedNode(6)
	node7 := NewSinglyLinkedNode(7)
	node8 := NewSinglyLinkedNode(8)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	list.InsertNodeAfterAt(node4, node6)
	list.String()
	assert.Equal(t, "1-->2-->3-->4-->6-->5-->nil", list.GetString(), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node7, node8), "InsertNodeAfterAt is incorrect")

	list = NewSinglyLinkedList(nil)
	node1 = NewSinglyLinkedNode(1)
	node2 = NewSinglyLinkedNode(2)
	node3 = NewSinglyLinkedNode(3)
	node4 = NewSinglyLinkedNode(4)
	node5 = NewSinglyLinkedNode(5)
	node6 = NewSinglyLinkedNode(6)
	node7 = NewSinglyLinkedNode(7)
	node8 = NewSinglyLinkedNode(8)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)
	list.InsertNodeAfterValueAt(3, node6)
	list.String()

	assert.Equal(t, "1-->2-->3-->6-->4-->5-->nil", list.GetString(), "InsertNodeAfterValueAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node7, node8), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, node6, list.GetNodeAtIndex(3), "GetNodeAtIndex is incorrect")

	list = NewSinglyLinkedList(nil)
	node1 = NewSinglyLinkedNode(1)
	node2 = NewSinglyLinkedNode(2)
	node3 = NewSinglyLinkedNode(3)
	node4 = NewSinglyLinkedNode(4)
	node5 = NewSinglyLinkedNode(5)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	ret := list.DeleteNode(node3)
	list.String()
	assert.Equal(t, "1-->2-->4-->5-->nil", list.GetString(), "DeleteNode is incorrect")
	assert.Equal(t, true, ret, "DeleteNode is incorrect")
	list.DeleteNode(node1)
	list.String()
	assert.Equal(t, "2-->4-->5-->nil", list.GetString(), "DeleteNode is incorrect")
	list.DeleteNode(node5)
	list.String()
	assert.Equal(t, "2-->4-->nil", list.GetString(), "DeleteNode is incorrect")
	ret = list.DeleteNode(node5)
	assert.Equal(t, false, ret, "DeleteNode is incorrect")

	list = NewSinglyLinkedList(nil)
	node1 = NewSinglyLinkedNode(1)
	node2 = NewSinglyLinkedNode(2)
	node3 = NewSinglyLinkedNode(3)
	node4 = NewSinglyLinkedNode(4)
	node5 = NewSinglyLinkedNode(5)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	list.String()
	list.DeleteNodeAtIndex(3)
	assert.Equal(t, "1-->2-->3-->5-->nil", list.GetString(), "DeleteNodeAtIndex is incorrect")
	list.DeleteNodeAtIndex(0)
	list.String()
	assert.Equal(t, "2-->3-->5-->nil", list.GetString(), "DeleteNodeAtIndex is incorrect")
	list.DeleteNodeAtIndex(list.length - 1)
	list.String()
	assert.Equal(t, "2-->3-->nil", list.GetString(), "DeleteNodeAtIndex is incorrect")

}
