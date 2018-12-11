package linkedlist

import (
	"fmt"
)

// SinglyLinkedNode 单向链表节点
type SinglyLinkedNode struct {
	next  *SinglyLinkedNode
	value interface{}
}

// NewSinglyLinkedNode 构造函数
func NewSinglyLinkedNode(v interface{}) *SinglyLinkedNode {
	return &SinglyLinkedNode{value: v, next: nil}
}

// GetValue 获取节点的内容
func (node *SinglyLinkedNode) GetValue() interface{} {
	return node.value
}

// GetNext 获取节点的下一个节点
func (node *SinglyLinkedNode) GetNext() *SinglyLinkedNode {
	return node.next
}

// SinglyLinkedList 单向链表, 带哨兵节点(空头节点)
type SinglyLinkedList struct {
	head   *SinglyLinkedNode
	length int64
}

// NewSinglyLinkedList 单向链表构造函数
func NewSinglyLinkedList(node *SinglyLinkedNode) *SinglyLinkedList {
	head := NewSinglyLinkedNode(nil)
	head.next = node
	return &SinglyLinkedList{head: head, length: 0}
}

// NewSinglyLinkedListFromNode 从一个头节点中构建出一个链表对象
func NewSinglyLinkedListFromNode(node *SinglyLinkedNode) *SinglyLinkedList {
	if node == nil {
		panic("node can't be nil!")
	}

	list := NewSinglyLinkedList(nil)
	list.InsertNode(node)

	return list
}

// GetFirstNode 返回第一个数据节点
func (list *SinglyLinkedList) GetFirstNode() *SinglyLinkedNode {
	return list.head.next
}

// InsertNode 在链表末尾插入一个节点
func (list *SinglyLinkedList) InsertNode(node *SinglyLinkedNode) {
	p := list.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
	list.length++
}

// InsertNodeHead 在链表头部插入一个节点
func (list *SinglyLinkedList) InsertNodeHead(node *SinglyLinkedNode) {
	p := list.head
	node.next = p.next
	p.next = node
	list.length++
}

// InsertNodeAfterAt 在链表dest节点后插入一个新节点
// @return 如果指定节点不存在则插入失败, 返回 false
func (list *SinglyLinkedList) InsertNodeAfterAt(dest *SinglyLinkedNode, node *SinglyLinkedNode) (inserted bool) {
	p := list.head
	inserted = false
	for p.next != nil && p.next != dest {
		p = p.next
	}
	// 说明找到指定节点了, 开始插入
	if p.next == dest {
		node.next = p.next.next
		p.next.next = node
		inserted = true
		list.length++
	}
	return
}

// InsertNodeAfterValueAt 在链表dest节点后插入一个新节点
func (list *SinglyLinkedList) InsertNodeAfterValueAt(dest interface{}, node *SinglyLinkedNode) (inserted bool) {
	p := list.head
	inserted = false
	for p.next != nil && p.next.value != dest {
		p = p.next
	}
	// 说明找到指定节点了, 开始插入
	if p.next.value == dest {
		node.next = p.next.next
		p.next.next = node
		inserted = true
		list.length++
	}
	return
}

// GetNodeAtIndex 获取指定位置的节点
func (list *SinglyLinkedList) GetNodeAtIndex(index int64) *SinglyLinkedNode {
	if index > list.length-1 || index < 0 {
		panic("out of range")
	}

	p := list.head
	for index > -1 {
		p = p.next
		index--
	}

	return p
}

// DeleteNode 删除指定节点
func (list *SinglyLinkedList) DeleteNode(dest *SinglyLinkedNode) (deleted bool) {
	if dest == nil {
		panic("dest can't be nil!")
	}

	p := list.head
	for p.next != nil && p.next != dest {
		p = p.next
	}
	if p.next == dest {
		p.next = p.next.next
		// 这里golang会自动delete dest
		deleted = true
		list.length--
	}
	return
}

// DeleteNodeAtIndex 删除指定位置的节点
func (list *SinglyLinkedList) DeleteNodeAtIndex(index int64) {
	if index > list.length-1 || index < 0 {
		panic("out of range")
	}

	p := list.head

	// 让p指向目标节点的前一个节点, 即 p.next = list[index]
	for index > 0 {
		p = p.next
		index--
	}
	// 因为事先已经对长度进行判断了, 所以这里p.next可以确定是非nil
	p.next = p.next.next
	list.length--
	return
}

// 打印出节点
func (list *SinglyLinkedList) String() {
	p := list.head
	for p.next != nil {
		fmt.Printf("%v-->", p.next.value)
		p = p.next
	}
	fmt.Printf("nil\n")
}

// GetString 获取整条链, 返回一个字符串
func (list *SinglyLinkedList) GetString() (listString string) {
	p := list.head
	for p.next != nil {
		listString += fmt.Sprintf("%v-->", p.next.value)
		p = p.next
	}
	listString += fmt.Sprintf("nil")
	return
}
