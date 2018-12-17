package stack

import "fmt"

/*
///顺序栈 使用数组实现, 因为栈比较简单, 所以就把两种存储结构都实现一下吧
*/

//Stack 栈
//其实在golang里面, slice本质就是一个动态扩容数组, 所以线型栈用slice实理论上也是无限容量
type Stack struct {
	items []interface{} //元素
	count int           //栈中存在的元素个数
	cap   int           //栈最大容量
}

//NewStack 构造函数
func NewStack(cap int) *Stack {
	return &Stack{items: make([]interface{}, cap), count: 0, cap: cap}
}

// Count 获取栈中元素数量
func (stack *Stack) Count() int {
	return stack.count
}

// Cap 获取栈最大容量
func (stack *Stack) Cap() int {
	return stack.cap
}

// Push 压栈, 当栈满的时候返回false
func (stack *Stack) Push(item interface{}) (success bool) {
	if stack.count == stack.cap {
		return false
	}
	stack.count++
	stack.items[stack.count-1] = item
	return true
}

// Pop 出栈, 栈为空时返回nil
func (stack *Stack) Pop() interface{} {
	if stack.count == 0 {
		return nil
	}

	popItem := stack.items[stack.count-1]
	// 释放引用, 避免内存泄漏
	stack.items[stack.count-1] = nil
	stack.items = stack.items[:stack.count-1]
	stack.count--
	return popItem
}

func (stack *Stack) String() string {
	var stackString string
	for index := 0; index < stack.count; index++ {
		stackString += fmt.Sprintf("%v-", stack.items[index])
	}
	return stackString
}
