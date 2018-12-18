package queue

import "fmt"

// Queue 顺序队列
// 这里采用循环队列, 因为循环队列是最常见也是最高效的队列实现方案
type Queue struct {
	q     []interface{}
	count int // 队中元素数量
	cap   int // 队列容量
	front int // 队头下标
	rear  int // rear-1 是队尾元素下标
}

// NewQueue 构造函数
func NewQueue(cap int) *Queue {
	// 因为队列有一个空的占位符, 所以总需要多申请一个空间来确保实际有效容量有cap个
	return &Queue{q: make([]interface{}, cap+1, cap+1), cap: cap + 1}
}

// Cap 队列容量
func (q *Queue) Cap() int {
	return q.cap - 1
}

// Count 队列元素数量
func (q *Queue) Count() int {
	return q.count
}

// IsEmpty 队列是否为空
// 这是通用的队空判断条件, 当前实现也可以直接通过比较count, cap判断. count==0时空, count==cap时满
func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull 队列是否为空
// 这是通用的队满判断条件, 当前实现也可以直接通过比较count, cap判断. count==0时空, count==cap时满
func (q *Queue) IsFull() bool {
	return (q.rear+1)%q.cap == q.front
}

// EnQueue 入队, 队满时返回bool
func (q *Queue) EnQueue(ele interface{}) bool {
	if ele == nil {
		panic("ele can't be nil")
	}

	if q.IsFull() {
		return false
	}
	q.q[q.rear] = ele
	q.rear = (q.rear + 1) % q.cap
	q.count++
	return true
}

// DeQueue 出队, 对空时返回nil
func (q *Queue) DeQueue() interface{} {

	if q.IsEmpty() {
		return nil
	}
	ele := q.q[q.front]
	q.count--
	q.front = (q.front + 1) % q.cap
	return ele
}

func (q *Queue) String() string {
	if q.IsEmpty() {
		return "-"
	}

	var qString string
	i := q.front
	for i != q.rear {
		qString += fmt.Sprintf("%v-", q.q[i])
		i = (i + 1) % q.cap
	}
	return qString
}
