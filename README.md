# DSA
 Data Structure &amp; Algorithm, 算法之美, Go语言实现

# 实现必修的数据结构与算法

## [单向链表](https://github.com/cocos543/DSA/blob/master/core/linkedlist/singly_linked_list.go)

### 链表常见操作
* 反转单向链表 ✓

    时间复杂度O(n), 空间复杂度O(1)

* 链表中环的检测 ✓

    时间复杂度O(n), 空间复杂度O(1)

* 合并两个有序链表 ✓

    时间复杂度O(n), 空间复杂度O(1)

* 删除链表倒数第N个节点 ✓

    时间复杂度O(n), 空间复杂度O(1)

* 求链表的中间节点 ✓

    时间复杂度O(n), 空间复杂度O(1)

## [双向链表](https://github.com/cocos543/DSA/blob/master/core/linkedlist/double_linked_lists.go)

## [栈](https://github.com/cocos543/DSA/blob/master/core/stack)

### 栈存储结构
* [顺序栈](https://github.com/cocos543/DSA/blob/master/core/stack/stack_sequential_storage.go) ✓

* [链式栈](https://github.com/cocos543/DSA/blob/master/core/stack/stack_linked_storage.go) ✓

## [队列](https://github.com/cocos543/DSA/blob/master/core/queue)

* [顺序队列](https://github.com/cocos543/DSA/blob/master/core/queue/queue.go) ✓

## [排序](https://github.com/cocos543/DSA/blob/master/core/sort)

### 稳定的排序算法

* [冒泡排序](https://github.com/cocos543/DSA/blob/master/core/sort/sort.go) ✓

    最好时间复杂度O(n), 最坏时间复杂度O(n^2), 平均时间复杂度O(n^2), 空间复杂度O(1)
![image](https://github.com/cocos543/DSA/blob/master/resource/img/BubbleSort.jpg)

* [直接插入排序](https://github.com/cocos543/DSA/blob/master/core/sort/sort.go) ✓

    最好时间复杂度O(n), 最坏时间复杂度O(n^2), 平均时间复杂度O(n^2), 空间复杂度O(1)
![image](https://github.com/cocos543/DSA/blob/master/resource/img/StraightInsertionSort.jpg)

* [计数排序(特殊桶排序)](https://github.com/cocos543/DSA/blob/master/core/sort/sort.go) ✓

    时间复杂度O(n), 空间复杂度O(n)

* [归并排序](https://github.com/cocos543/DSA/blob/master/core/sort/sort_merging.go) ✓

    最好时间复杂度O(nlogn), 最坏时间复杂度O(nlogn), 平均时间复杂度O(nlogn), 空间复杂度O(n)
![image](https://github.com/cocos543/DSA/blob/master/resource/img/MergingSort.jpg)

### 不稳定的排序算法

* [快速排序](https://github.com/cocos543/DSA/blob/master/core/sort/sort_quick.go) ✓

    最好时间复杂度O(nlogn), 最坏时间复杂度O(n^2), 平均时间复杂度O(nlogn), 空间复杂度O(1)
![image](https://github.com/cocos543/DSA/blob/master/resource/img/QuickSort.jpg)

* [快排思想查找第k大元素](https://github.com/cocos543/DSA/blob/master/core/sort/sort_quick.go) ✓

    时间复杂度O(n), 空间复杂度O(1)

    分区函数算法示意图
![image](https://github.com/cocos543/DSA/blob/master/resource/img/QuickSortPartition.jpg)

## [查找](https://github.com/cocos543/DSA/blob/master/core/search/search)

### 二分查找

* [二分查找](https://github.com/cocos543/DSA/blob/master/core/search/search.go) ✓

   时间复杂度O(logn), 空间复杂度O(1)
![image](https://github.com/cocos543/DSA/blob/master/resource/img/BinarySearch.jpg)

* [求解平方根](https://github.com/cocos543/DSA/blob/master/core/search/search.go) ✓

   时间复杂度O(logn), 空间复杂度O(1)

### [四种二分查找的变形算法](https://github.com/cocos543/DSA/blob/master/core/search/search.go)

1. 查找第一个值等于给定值的元素 ✓

2. 查找最后一个值等于给定值的元素 ✓

3. 查找第一个大于等于给定值的元素 ✓

4. 查找最后一个小于等于给定值的元素 ✓