package linkedlist

import "fmt"

// DoubleLinkedNode 双向链表节点
type DoubleLinkedNode struct {
	value interface{}
	perv  *DoubleLinkedNode
	next  *DoubleLinkedNode
}

// NewDoubleLinkedNode 构造函数
func NewDoubleLinkedNode(v interface{}) *DoubleLinkedNode {
	return &DoubleLinkedNode{value: v}
}

// Value 获取节点的内容
func (node *DoubleLinkedNode) Value() interface{} {
	return node.value
}

// Next 获取节点的下一个节点
func (node *DoubleLinkedNode) Next() *DoubleLinkedNode {
	return node.next
}

// Perv 获取节点的下一个节点
func (node *DoubleLinkedNode) Perv() *DoubleLinkedNode {
	return node.perv
}

// DoubleLinkedLists 双向链表, 带哨兵节点(空头节点)
type DoubleLinkedLists struct {
	head   *DoubleLinkedNode
	length int
}

// NewDoubleLinkedLists 双向链表构造函数
func NewDoubleLinkedLists() *DoubleLinkedLists {
	head := NewDoubleLinkedNode(nil)
	return &DoubleLinkedLists{head: head, length: 0}
}

// NewDoubleLinkedListsFromNode 从一个头节点中构建出一个链表对象
func NewDoubleLinkedListsFromNode(node *DoubleLinkedNode) *DoubleLinkedLists {
	if node == nil {
		panic("node can't be nil!")
	}

	list := NewDoubleLinkedLists()
	list.InsertNode(node)

	return list
}

// Length 获取链表长度
func (list *DoubleLinkedLists) Length() int {
	return list.length
}

// GetFirstNode 返回第一个数据节点
func (list *DoubleLinkedLists) GetFirstNode() *DoubleLinkedNode {
	return list.head.next
}

// InsertNode 在链表末尾插入一个节点
func (list *DoubleLinkedLists) InsertNode(node *DoubleLinkedNode) {
	p := list.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
	node.perv = p
	list.length++
}

// InsertNodeHead 在链表头部插入一个节点
func (list *DoubleLinkedLists) InsertNodeHead(node *DoubleLinkedNode) {
	p := list.head
	if p.next != nil {
		p.next.perv = node
	}
	node.next = p.next

	p.next = node
	node.perv = p

	list.length++
}

// InsertNodeAfterAt 在链表dest节点后插入一个新节点
// @return 如果指定节点不存在则插入失败, 返回 false
func (list *DoubleLinkedLists) InsertNodeAfterAt(dest *DoubleLinkedNode, node *DoubleLinkedNode) (inserted bool) {
	p := list.head
	inserted = false
	for p.next != nil && p.next != dest {
		p = p.next
	}
	// 说明找到指定节点了, 开始插入
	if p.next == dest {
		if p.next.next != nil {
			p.next.next.perv = node
		}
		node.next = p.next.next

		p.next.next = node
		node.perv = p.next
		inserted = true
		list.length++
	}
	return
}

// InsertNodeAfterValueAt 在链表dest节点后插入一个新节点
func (list *DoubleLinkedLists) InsertNodeAfterValueAt(dest interface{}, node *DoubleLinkedNode) (inserted bool) {
	p := list.head
	inserted = false
	for p.next != nil && p.next.value != dest {
		p = p.next
	}
	// 说明找到指定节点了, 开始插入
	if p.next.value == dest {
		if p.next.next != nil {
			p.next.next.perv = node
		}
		node.next = p.next.next

		p.next.next = node
		node.perv = p.next
		inserted = true
		list.length++
	}
	return
}

// GetNodeAtIndex 获取指定位置的节点
func (list *DoubleLinkedLists) GetNodeAtIndex(index int) *DoubleLinkedNode {
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
func (list *DoubleLinkedLists) DeleteNode(dest *DoubleLinkedNode) (deleted bool) {
	if dest == nil {
		panic("dest can't be nil!")
	}

	p := list.head
	for p.next != nil && p.next != dest {
		p = p.next
	}
	if p.next == dest {
		if p.next.next != nil {
			p.next.next.perv = p
		}
		p.next = p.next.next
		// 这里golang会自动delete dest
		deleted = true
		list.length--
	}
	return
}

// DeleteNodeAtIndex 删除指定位置的节点
func (list *DoubleLinkedLists) DeleteNodeAtIndex(index int) {
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
	if p.next.next != nil {
		p.next.next.perv = p
	}

	p.next = p.next.next

	list.length--
	return
}

func (list *DoubleLinkedLists) String() string {
	var listString1 string
	var listString2 string
	p := list.head
	for p.next != nil {
		listString1 += fmt.Sprintf("%v-->", p.next.value)
		p = p.next
	}
	listString1 += fmt.Sprintf("nil")
	listString1 += "\n"
	for p != list.head {
		listString2 = fmt.Sprintf("<--%v", p.value) + listString2
		p = p.perv
	}

	listString2 = fmt.Sprintf("head") + listString2
	return listString1 + listString2 + "\n"
}
