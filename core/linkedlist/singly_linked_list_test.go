package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertList(t *testing.T) {
	list := NewSinglyLinkedList()
	for i := 0; i < 10; i++ {
		node := NewSinglyLinkedNode(i + 1)
		list.InsertNode(node)
	}
	fmt.Println(list)
	assert.Equal(t, 1, list.head.next.Value(), "InsertNode is incorrect")

	list = NewSinglyLinkedList()
	for i := 0; i < 10; i++ {
		node := NewSinglyLinkedNode(i + 1)
		list.InsertNodeHead(node)
	}
	fmt.Println(list)

	list = NewSinglyLinkedList()
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
	fmt.Println(list)
	assert.Equal(t, "1-->2-->3-->4-->6-->5-->nil", list.String(), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node7, node8), "InsertNodeAfterAt is incorrect")

	list = NewSinglyLinkedList()
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
	fmt.Println(list)

	assert.Equal(t, "1-->2-->3-->6-->4-->5-->nil", list.String(), "InsertNodeAfterValueAt is incorrect")
	assert.Equal(t, false, list.InsertNodeAfterAt(node7, node8), "InsertNodeAfterAt is incorrect")
	assert.Equal(t, node6, list.GetNodeAtIndex(3), "GetNodeAtIndex is incorrect")

	list = NewSinglyLinkedList()
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
	fmt.Println(list)
	assert.Equal(t, "1-->2-->4-->5-->nil", list.String(), "DeleteNode is incorrect")
	assert.Equal(t, true, ret, "DeleteNode is incorrect")
	list.DeleteNode(node1)
	fmt.Println(list)
	assert.Equal(t, "2-->4-->5-->nil", list.String(), "DeleteNode is incorrect")
	list.DeleteNode(node5)
	fmt.Println(list)
	assert.Equal(t, "2-->4-->nil", list.String(), "DeleteNode is incorrect")
	ret = list.DeleteNode(node5)
	assert.Equal(t, false, ret, "DeleteNode is incorrect")

	list = NewSinglyLinkedList()
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

	fmt.Println(list)
	list.DeleteNodeAtIndex(3)
	assert.Equal(t, "1-->2-->3-->5-->nil", list.String(), "DeleteNodeAtIndex is incorrect")
	list.DeleteNodeAtIndex(0)
	fmt.Println(list)
	assert.Equal(t, "2-->3-->5-->nil", list.String(), "DeleteNodeAtIndex is incorrect")
	list.DeleteNodeAtIndex(list.length - 1)
	fmt.Println(list)
	assert.Equal(t, "2-->3-->nil", list.String(), "DeleteNodeAtIndex is incorrect")

}

func TestReverseListNode(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	node3 := NewSinglyLinkedNode(3)
	node4 := NewSinglyLinkedNode(4)
	node5 := NewSinglyLinkedNode(5)
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)
	fmt.Println(list)
	t.Log("Reversed:")
	head := ReverseList(list.GetFirstNode())
	list = NewSinglyLinkedListFromNode(head)
	fmt.Println(list)
	assert.Equal(t, "5-->4-->3-->2-->1-->nil", list.String(), "ReverseList is incorrect")
}

func TestPalindrome(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewSinglyLinkedNode("a")
	list.InsertNode(node1)
	assert.Equal(t, true, IsPalindrome(list.GetFirstNode()), "IsPalindrome is incorrect")

	list = NewSinglyLinkedList()
	node1 = NewSinglyLinkedNode("a")
	node2 := NewSinglyLinkedNode("a")
	list.InsertNode(node1)
	assert.Equal(t, true, IsPalindrome(list.GetFirstNode()), "IsPalindrome is incorrect")

	list = NewSinglyLinkedList()
	node1 = NewSinglyLinkedNode("a")
	node2 = NewSinglyLinkedNode("b")
	node3 := NewSinglyLinkedNode("c")
	node4 := NewSinglyLinkedNode("d")
	node5 := NewSinglyLinkedNode("c")
	node6 := NewSinglyLinkedNode("b")
	node7 := NewSinglyLinkedNode("a")
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)
	list.InsertNode(node6)
	list.InsertNode(node7)
	fmt.Println(list)

	assert.Equal(t, true, IsPalindrome(list.GetFirstNode()), "IsPalindrome is incorrect")
	assert.Equal(t, "a-->b-->c-->d-->c-->b-->a-->nil", list.String(), "IsPalindrome is incorrect")

	list = NewSinglyLinkedList()
	node1 = NewSinglyLinkedNode("a")
	node2 = NewSinglyLinkedNode("b")
	node3 = NewSinglyLinkedNode("c")
	node4 = NewSinglyLinkedNode("c")
	node5 = NewSinglyLinkedNode("b")
	node6 = NewSinglyLinkedNode("a")
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)
	list.InsertNode(node6)
	fmt.Println(list)

	assert.Equal(t, true, IsPalindrome(list.GetFirstNode()), "IsPalindrome is incorrect")
	assert.Equal(t, "a-->b-->c-->c-->b-->a-->nil", list.String(), "IsPalindrome is incorrect")
}

func TestIsLoopLinkedList(t *testing.T) {
	list := NewSinglyLinkedList()
	node1 := NewSinglyLinkedNode(1)
	list.InsertNode(node1)

	assert.Equal(t, false, IsLoopLinkedList(list.GetFirstNode()), "IsLoopLinkedList is incorrect")
	node1.next = node1
	assert.Equal(t, true, IsLoopLinkedList(list.GetFirstNode()), "IsLoopLinkedList is incorrect")

	list = NewSinglyLinkedList()
	node1 = NewSinglyLinkedNode(1)
	node2 := NewSinglyLinkedNode(2)
	list.InsertNode(node1)
	list.InsertNode(node2)
	node2.next = node1

	assert.Equal(t, true, IsLoopLinkedList(list.GetFirstNode()), "IsLoopLinkedList is incorrect")

	list = NewSinglyLinkedList()
	node1 = NewSinglyLinkedNode(1)
	node2 = NewSinglyLinkedNode(2)
	node3 := NewSinglyLinkedNode(3)
	node4 := NewSinglyLinkedNode(4)
	node5 := NewSinglyLinkedNode(5)
	node5.next = node2
	list.InsertNode(node1)
	list.InsertNode(node2)
	list.InsertNode(node3)
	list.InsertNode(node4)
	list.InsertNode(node5)

	assert.Equal(t, true, IsLoopLinkedList(list.GetFirstNode()), "IsLoopLinkedList is incorrect")

}

func TestMergeTowOrderedList(t *testing.T) {
	listA := NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))
	listA.InsertNode(NewSinglyLinkedNode(2))

	listB := NewSinglyLinkedList()
	listB.InsertNode(NewSinglyLinkedNode(1))

	list := NewSinglyLinkedList()
	list.InsertNode(MergeTowOrderedList(listA.GetFirstNode(), listB.GetFirstNode()))
	fmt.Println(list)
	assert.Equal(t, "1-->1-->2-->nil", list.String(), "MergeTowOrderedList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(3))
	listA.InsertNode(NewSinglyLinkedNode(4))

	listB = NewSinglyLinkedList()
	listB.InsertNode(NewSinglyLinkedNode(1))
	listB.InsertNode(NewSinglyLinkedNode(3))
	listB.InsertNode(NewSinglyLinkedNode(4))
	listB.InsertNode(NewSinglyLinkedNode(5))

	list = NewSinglyLinkedList()
	list.InsertNode(MergeTowOrderedList(listB.GetFirstNode(), listA.GetFirstNode()))
	fmt.Println(list)
	assert.Equal(t, "1-->1-->2-->3-->3-->4-->4-->5-->nil", list.String(), "MergeTowOrderedList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(5))
	listA.InsertNode(NewSinglyLinkedNode(1))

	listB = NewSinglyLinkedList()
	listB.InsertNode(NewSinglyLinkedNode(1))

	list = NewSinglyLinkedList()
	list.InsertNode(MergeTowOrderedList(listA.GetFirstNode(), listB.GetFirstNode()))
	fmt.Println(list)
	assert.Equal(t, "5-->1-->1-->nil", list.String(), "MergeTowOrderedList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(4))
	listA.InsertNode(NewSinglyLinkedNode(3))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(1))

	listB = NewSinglyLinkedList()
	listB.InsertNode(NewSinglyLinkedNode(5))
	listB.InsertNode(NewSinglyLinkedNode(4))
	listB.InsertNode(NewSinglyLinkedNode(3))
	listB.InsertNode(NewSinglyLinkedNode(2))

	list = NewSinglyLinkedList()
	list.InsertNode(MergeTowOrderedList(listB.GetFirstNode(), listA.GetFirstNode()))
	fmt.Println(list)
	assert.Equal(t, "5-->4-->4-->3-->3-->2-->2-->1-->nil", list.String(), "MergeTowOrderedList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(-10))
	listA.InsertNode(NewSinglyLinkedNode(-6))
	listA.InsertNode(NewSinglyLinkedNode(-1))
	listA.InsertNode(NewSinglyLinkedNode(-1))
	listA.InsertNode(NewSinglyLinkedNode(-1))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(7))
	listA.InsertNode(NewSinglyLinkedNode(8))
	listA.InsertNode(NewSinglyLinkedNode(8))

	listB = NewSinglyLinkedList()
	listB.InsertNode(NewSinglyLinkedNode(-10))
	listB.InsertNode(NewSinglyLinkedNode(-10))
	listB.InsertNode(NewSinglyLinkedNode(-8))
	listB.InsertNode(NewSinglyLinkedNode(-6))
	listB.InsertNode(NewSinglyLinkedNode(0))
	listB.InsertNode(NewSinglyLinkedNode(2))
	listB.InsertNode(NewSinglyLinkedNode(5))
	listB.InsertNode(NewSinglyLinkedNode(7))

	list = NewSinglyLinkedList()
	list.InsertNode(MergeTowOrderedList(listB.GetFirstNode(), listA.GetFirstNode()))
	fmt.Println(list)
	assert.Equal(t, "-10-->-10-->-10-->-8-->-6-->-6-->-1-->-1-->-1-->0-->2-->2-->5-->7-->7-->8-->8-->nil", list.String(), "MergeTowOrderedList is incorrect")

}

func TestRemoveNthNodeFromEndofList(t *testing.T) {
	listA := NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))

	list := NewSinglyLinkedList()
	list.InsertNode(RemoveNthNodeFromEndofList(listA.GetFirstNode(), 1))
	fmt.Println(list)
	assert.Equal(t, "nil", list.String(), "RemoveNthNodeFromEndofList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(3))
	listA.InsertNode(NewSinglyLinkedNode(4))
	listA.InsertNode(NewSinglyLinkedNode(5))

	list = NewSinglyLinkedList()
	list.InsertNode(RemoveNthNodeFromEndofList(listA.GetFirstNode(), 1))

	fmt.Println(list)
	assert.Equal(t, "1-->2-->3-->4-->nil", list.String(), "RemoveNthNodeFromEndofList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(3))
	listA.InsertNode(NewSinglyLinkedNode(4))
	listA.InsertNode(NewSinglyLinkedNode(5))

	list = NewSinglyLinkedList()
	list.InsertNode(RemoveNthNodeFromEndofList(listA.GetFirstNode(), 4))

	fmt.Println(list)
	assert.Equal(t, "1-->3-->4-->5-->nil", list.String(), "RemoveNthNodeFromEndofList is incorrect")

	listA = NewSinglyLinkedList()
	listA.InsertNode(NewSinglyLinkedNode(1))
	listA.InsertNode(NewSinglyLinkedNode(2))
	listA.InsertNode(NewSinglyLinkedNode(3))
	listA.InsertNode(NewSinglyLinkedNode(4))
	listA.InsertNode(NewSinglyLinkedNode(5))

	list = NewSinglyLinkedList()
	list.InsertNode(RemoveNthNodeFromEndofList(listA.GetFirstNode(), 5))

	fmt.Println(list)
	assert.Equal(t, "2-->3-->4-->5-->nil", list.String(), "RemoveNthNodeFromEndofList is incorrect")

	fmt.Println(list)
}
