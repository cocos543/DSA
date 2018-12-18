package stack

import (
	"DSA/core/linkedlist"
	"fmt"
)

// LikedStack 链式栈 使用双向链表实现
type LikedStack struct {
	items *linkedlist.DoubleLinkedLists //元素
	cap   int                           //栈最大容量
	top   *linkedlist.DoubleLinkedNode  //指向栈顶节点
}

// NewLikedStack 构造函数
func NewLikedStack(cap int) *LikedStack {
	return &LikedStack{items: linkedlist.NewDoubleLinkedLists(), cap: cap}
}

// Count 获取栈中元素数量
func (stack *LikedStack) Count() int {
	return stack.items.Length()
}

// Cap 获取栈最大容量
func (stack *LikedStack) Cap() int {
	return stack.cap
}

// Push 压栈, 当栈满的时候返回false
func (stack *LikedStack) Push(item interface{}) (success bool) {
	if stack.Count() == stack.cap {
		return false
	}
	stack.top = linkedlist.NewDoubleLinkedNode(item)
	stack.items.InsertNode(stack.top)

	return true
}

// Pop 出栈, 栈为空时返回nil
func (stack *LikedStack) Pop() interface{} {
	if stack.Count() == 0 {
		return nil
	}

	popItem := stack.top
	stack.top = stack.top.Perv()
	stack.items.DeleteNode(popItem)
	return popItem.Value()
}

func (stack *LikedStack) String() string {
	var stackString string
	top := stack.top
	for index := 0; index < stack.Count(); index++ {
		stackString = fmt.Sprintf("%v-", top.Value()) + stackString
		top = top.Perv()
	}
	return stackString
}
