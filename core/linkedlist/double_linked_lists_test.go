package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertDoubleLinkedLists(t *testing.T) {
	list := NewDoubleLinkedLists()
	for i := 0; i < 10; i++ {
		node := NewDoubleLinkedNode(i + 1)
		list.InsertNode(node)
	}
	fmt.Println(list)
	assert.Equal(t, "1-->2-->3-->4-->5-->6-->7-->8-->9-->10-->nil\nhead<--1<--2<--3<--4<--5<--6<--7<--8<--9<--10\n", list.String(), "InsertNode is incorrect")
	assert.EqualValues(t, 10, list.length, "InsertNode is incorrect")

	list = NewDoubleLinkedLists()
	for i := 10; i > 0; i-- {
		node := NewDoubleLinkedNode(i)
		list.InsertNodeHead(node)
	}
	fmt.Println(list)
	assert.Equal(t, "1-->2-->3-->4-->5-->6-->7-->8-->9-->10-->nil\nhead<--1<--2<--3<--4<--5<--6<--7<--8<--9<--10\n", list.String(), "InsertNodeHead is incorrect")
	assert.EqualValues(t, 10, list.length, "InsertNodeHead is incorrect")

	list = NewDoubleLinkedLists()
	node1 := NewDoubleLinkedNode(1)
	node2 := NewDoubleLinkedNode(2)
	node3 := NewDoubleLinkedNode(3)
	node4 := NewDoubleLinkedNode(4)
	node5 := NewDoubleLinkedNode(5)
	node6 := NewDoubleLinkedNode(6)
	node7 := NewDoubleLinkedNode(7)
	node8 := NewDoubleLinkedNode(8)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	list.InsertNodeAfterAt(node4, node6)
	list.InsertNodeAfterAt(node1, node7)

	fmt.Println(list)
	assert.Equal(t, "1-->7-->2-->3-->4-->6-->5-->nil\nhead<--1<--7<--2<--3<--4<--6<--5\n", list.String(), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node8, node1), "InsertNodeAfterAt is incorrect")
	assert.EqualValues(t, 7, list.length, "InsertNodeAfterAt is incorrect")

	list = NewDoubleLinkedLists()
	node1 = NewDoubleLinkedNode(1)
	node2 = NewDoubleLinkedNode(2)
	node3 = NewDoubleLinkedNode(3)
	node4 = NewDoubleLinkedNode(4)
	node5 = NewDoubleLinkedNode(5)
	node6 = NewDoubleLinkedNode(6)
	node7 = NewDoubleLinkedNode(7)
	node8 = NewDoubleLinkedNode(8)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	list.InsertNodeAfterValueAt(4, node6)
	list.InsertNodeAfterValueAt(1, node7)

	fmt.Println(list)
	assert.Equal(t, "1-->7-->2-->3-->4-->6-->5-->nil\nhead<--1<--7<--2<--3<--4<--6<--5\n", list.String(), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node8, node1), "InsertNodeAfterAt is incorrect")

	assert.EqualValues(t, 7, list.length, "InsertNodeAfterValueAt is incorrect")
	assert.Equal(t, node1, list.GetNodeAtIndex(0), "GetNodeAtIndex is incorrect")
	assert.Equal(t, node5, list.GetNodeAtIndex(6), "GetNodeAtIndex is incorrect")

	list = NewDoubleLinkedLists()
	node1 = NewDoubleLinkedNode(1)
	node2 = NewDoubleLinkedNode(2)
	node3 = NewDoubleLinkedNode(3)
	node4 = NewDoubleLinkedNode(4)
	node5 = NewDoubleLinkedNode(5)
	node6 = NewDoubleLinkedNode(6)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)
	list.InsertNode(node6)

	list.DeleteNode(node1)
	ret := list.DeleteNode(node6)
	ret2 := list.DeleteNode(NewDoubleLinkedNode(10))
	fmt.Println(list)
	assert.Equal(t, true, ret, "DeleteNode is incorrect")
	assert.Equal(t, false, ret2, "DeleteNode is incorrect")
	assert.Equal(t, "2-->3-->4-->5-->nil\nhead<--2<--3<--4<--5\n", list.String(), "DeleteNode is incorrect")

	list.DeleteNodeAtIndex(0)
	list.DeleteNodeAtIndex(2)
	fmt.Println(list)
	assert.Equal(t, "3-->4-->nil\nhead<--3<--4\n", list.String(), "DeleteNode is incorrect")

}
