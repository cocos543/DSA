package sort

import "fmt"

// QuickSort 快速排序
// 最好时间复杂度O(nlogn), 最坏时间复杂度O(n^2), 平均时间复杂度O(nlogn), 空间复杂度O(1)
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func QuickSort(arr []interface{}, compare CompareF) {
	fmt.Printf("arr[%v,%v]\n", 0, len(arr)-1)
	quickSortGo(arr, 0, len(arr)-1, compare)
}

// quickSortGo 快速排序的算法部分
func quickSortGo(arr []interface{}, p, r int, compare CompareF) {
	if p >= r {
		fmt.Println()
		return
	}

	q := partition(arr, p, r, compare)

	fmt.Printf("Pivot:arr[%v]:%v\n", q, arr[q])
	fmt.Printf("%v\n\n", arr)

	fmt.Printf("arr[%v,%v]\n", p, q-1)
	quickSortGo(arr, p, q-1, compare)
	fmt.Printf("arr[%v,%v]\n", q+1, r)
	quickSortGo(arr, q+1, r, compare)
}

// partition 分区函数. 对元素进行分区, 将大于等于(小于等于)arr[r]的元素, 全部放在arr[r]的左边, 小于(大于)arr[r]的元素放在右边
// @return 返回分区标志的下标
func partition(arr []interface{}, p, r int, compare CompareF) int {
	//注: 以从递增排序为例注释.
	//将最后一个元素作为分区标志
	pivot := r
	//i始终指向从左数起, 第一个大于arr[pivot]的元素下标
	//游标j从左向右移动, 如果遇到arr[j]小于arr[pivot]的, 则交换arr[j]和arr[pivot]的位置, 同时i++, 这样i会一直标记着从左数起, 第一个大于arr[pivot]的元素下标
	i := p
	for j := p; j < pivot; j++ {
		//由于这里存在元素交换, 所以快排是不稳定算法(因为你永远的不知道被交换的元素, 后面是不是也要相同的, 这样他们的相对顺序就可能发生变化了)
		if compare(arr[j], arr[pivot]) {
			Swap(arr, i, j)
			i++
		}
	}

	//最后交换arr[pivot]和arr[i]的位置, 则可以确保arr[pivot]左边的元素都比自己小, 右边的元素都比自己大
	Swap(arr, i, pivot)
	return i
}
