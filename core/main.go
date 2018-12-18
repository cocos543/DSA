package main

import (
	. "DSA/core/queue"
	"fmt"
)

func main() {
	fmt.Println("Hello, Data Structure & Algorithm.")

	queue := NewQueue(1)
	fmt.Println(queue.EnQueue(1))
}
