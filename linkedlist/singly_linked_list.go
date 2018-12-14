package linkedlist

import (
	"fmt"
)

// SinglyLinkedNode 单向链表节点
type SinglyLinkedNode struct {
	value interface{}
	next  *SinglyLinkedNode
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
func NewSinglyLinkedList() *SinglyLinkedList {
	head := NewSinglyLinkedNode(nil)
	return &SinglyLinkedList{head: head, length: 0}
}

// NewSinglyLinkedListFromNode 从一个头节点中构建出一个链表对象
func NewSinglyLinkedListFromNode(node *SinglyLinkedNode) *SinglyLinkedList {
	if node == nil {
		panic("node can't be nil!")
	}

	list := NewSinglyLinkedList()
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

func (list *SinglyLinkedList) String() string {
	var listString string
	p := list.head
	for p.next != nil {
		listString += fmt.Sprintf("%v-->", p.next.value)
		p = p.next
	}
	listString += fmt.Sprintf("nil")
	return listString
}

// GetMedianNode 获取链表的中间节点
// 算法: 采用快慢指针定位法
func GetMedianNode(node *SinglyLinkedNode) *SinglyLinkedNode {
	if node == nil {
		panic("node can't be nil!")
	}

	head := NewSinglyLinkedNode(nil)
	head.next = node

	// p1 每次前进1步, p2每次前进2步
	p1 := head
	p2 := head

	for p2 != nil && p2.next != nil {
		p1 = p1.next
		p2 = p2.next.next
	}
	return p1
}

// ReverseList 反转链表
/// node 为链表第一个节点
func ReverseList(node *SinglyLinkedNode) *SinglyLinkedNode {
	if node == nil {
		return nil
	}

	head := NewSinglyLinkedNode(nil)
	head.next = node

	// per用于保存cur的前一个节点, curNext用于保存cur的下一个节点
	var per *SinglyLinkedNode
	cur := head.next
	for cur != nil {
		curNext := cur.next
		cur.next = per
		per = cur
		cur = curNext
	}
	return per
}

// IsPalindrome 判断回文串.
// node 为链表第一个节点.
// 算法: 不引入数组结构, 直接在链上进行判断.
func IsPalindrome(node *SinglyLinkedNode) (isPalindrome bool) {
	if node == nil {
		panic("node can't be nil!")
	}

	head := NewSinglyLinkedNode(nil)
	head.next = node

	// pm指向中位节点, pm.next 为后部分数据的第一个节点
	var pm *SinglyLinkedNode
	pm = GetMedianNode(head.next)
	fmt.Println("Median node is ", pm.value)
	// 反转后半部分
	pm.next = ReverseList(pm.next)

	// 开始检查回文
	pa := head
	pb := pm

	isPalindrome = true
	for pb.next != nil {
		if pa.next.value != pb.next.value {
			isPalindrome = false
		}
		pa = pa.next
		pb = pb.next
	}

	// 还原现场
	pm.next = ReverseList(pm.next)

	return
}

// IsLoopLinkedList 判断链表中是否含有环
// 算法: 采用快慢指针识别.
// 证明方法:
//     假设B的速度是m, A的速度是n, m > n, m和n都是整数. 当B, A速度都下降n时, 则A速度0, B速度m-n, 即每次循环, B可以前进m-n, 如果链表有环,环长度为N,则
//     B需要 N / (m - n) 次循环可以跑回原点与A相遇. 令 m =2, n = 1, 则循环 N/1 = N次时, 如果P(B) = P(A), 则是循环单向链表. 否则, 则不是循环单向链表
//     同理, 当A, B先后进入环中, 即起点不同, 假设A速度移动速度降为0, B的速度为m-n, 可知只要存在环, 则B一定会再次遇上A.
func IsLoopLinkedList(node *SinglyLinkedNode) (isLoop bool) {
	if node == nil {
		panic("node can't be nil!")
	}
	head := NewSinglyLinkedNode(nil)
	head.next = node

	pa := head
	pb := head

	// do while
	for {
		pa = pa.next
		pb = pb.next.next
		if pb == nil || pb.next == nil || pb == pa {
			break
		}
	}

	if pb == pa {
		isLoop = true
	} else {
		isLoop = false
	}

	return
}

// MergeTowOrderedList 合并两个有序链表, 合并之后保持顺序
// 这里规定节点存放int型数据, 两个链表的顺序是相同的
// 算法: 采用哨兵简化算法, 将b链上的节点插入到a链中
func MergeTowOrderedList(nodeA *SinglyLinkedNode, nodeB *SinglyLinkedNode) *SinglyLinkedNode {

	if nodeA == nil || nodeB == nil {
		panic("node can't be nil!")
	}

	pa := nodeA
	pb := nodeB

	// 先确认是否为升序
	isAsc := false
	OrderConfirm := false
	for pa.next != nil {
		if pa.value.(int) < pa.next.value.(int) {
			isAsc = true
			OrderConfirm = true
			break
		} else if pa.value.(int) > pa.next.value.(int) {
			isAsc = false
			OrderConfirm = true
		}
		pa = pa.next
	}

	// 如果无法从第一条链确认顺序, 那说明第一条了的全部节点的值都是相同的, 只能从第二条链确认了.
	if !OrderConfirm {
		for pb.next != nil {
			if pb.value.(int) < pb.next.value.(int) {
				isAsc = true
				OrderConfirm = true
				break
			} else if pb.value.(int) > pb.next.value.(int) {
				isAsc = false
				OrderConfirm = true
			}
			pb = pb.next
		}
	}

	// 如果两条链都无法确认顺序, 那说明两条链中任意一条自身的所有节点值都相同(可能只有1个节点), 这个时候只需要简单把一条接到另一条后面即可
	if !OrderConfirm {
		if pa.value.(int) < pb.value.(int) {
			pa.next = nodeB
			return nodeA
		}
		pb.next = nodeA
		return nodeB
	}

	// 哨兵节点
	headA := NewSinglyLinkedNode(nil)
	headA.next = nodeA

	headB := NewSinglyLinkedNode(nil)
	headB.next = nodeB

	pa = headA
	pb = headB

	for pa.next != nil && pb.next != nil {
		if (isAsc && pa.next.value.(int) < pb.next.value.(int)) || (!isAsc && pa.next.value.(int) > pb.next.value.(int)) {
			pa = pa.next
			continue
		}
		// 将pb.next插入到pa.next左边
		pbNextNext := pb.next.next
		pb.next.next = pa.next
		pa.next = pb.next
		pb.next = pbNextNext

	}

	// 把剩余的部分接到长链上
	// 这里要区分某一条链只有1个节点的情况
	if pb.next != nil {
		// 当a链到达末尾时, b链还有剩余节点, 则b链剩余部分直接插入到a后
		pa.next = pb.next
	}

	return headA.next
}

// RemoveNthNodeFromEndofList 移除倒数第N个节点
// 算法: 带哨兵节点, 这里假设n一直都是合法的.
func RemoveNthNodeFromEndofList(node *SinglyLinkedNode, n int) *SinglyLinkedNode {
	head := NewSinglyLinkedNode(nil)
	head.next = node

	// pn 最终将指向倒数第n个节点的前一个节点
	pn := head
	plast := head
	for n != 0 {
		plast = plast.next
		n--
	}

	for plast.next != nil {
		pn = pn.next
		plast = plast.next
	}

	// 这里pn.next就是倒数第n个节点, 直接移除就可以
	fmt.Println("Delete ", pn.next.value)
	pn.next = pn.next.next
	return head.next
}
