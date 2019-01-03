package sort

import "fmt"

// QuickSort 快速排序, 不稳定的
// 最好时间复杂度O(nlogn), 最坏时间复杂度O(n^2), 平均时间复杂度O(nlogn), 空间复杂度O(1)
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func QuickSort(arr []interface{}, compare CompareF) {
	fmt.Printf("arr[%v,%v]\n", 0, len(arr)-1)
	quickSortGo(arr, 0, len(arr)-1, compare)
}

// quickSortGo 快速排序的算法部分
func quickSortGo(arr []interface{}, p, r int, compare CompareF) {
	// 当q=0时, 下一个分区是(0,-1), 所以出现p小于r的情况,说明分区只有一个元素了可以直接返回.
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
	//游标j从左向右移动, 如果遇到arr[j]小于arr[pivot]的, 则交换arr[j]和arr[i]的位置, 同时i++, 这样i会一直标记着从左数起, 第一个大于arr[pivot]的元素下标
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

// FindKthLargest 利用快排思想高效查到无序数组中第K大的元素, 时间复杂度O(n)
func FindKthLargest(arr []interface{}, k int) interface{} {
	return findKthLargestGO(arr, 0, len(arr)-1, k)
}

// 查找算法的真实部分
// p 数组第一个元素, r 数组最后一个元素, k 第k大
func findKthLargestGO(arr []interface{}, p, r, k int) interface{} {
	// 区间里只有1个元素, 则该元素就是第k大了
	if p >= r {
		return arr[p]
	}

	cmp := func(a, b interface{}) bool {
		return a.(int) >= b.(int)
	}
	//经过分区之后, q左边比arr[q]小, 右边的比arr[q]大, 所以q就是第q+1大
	q := partition(arr, p, r, cmp)
	if q+1 == k {
		return arr[q]
	}

	//k > q+1, 说明第k大的数在左边区间, 也就是arr[p, q]
	//则个时候需要在arr[p, q-1]区间里面找到第k大的数, 这个数就是数组的第k大
	if k < q+1 {
		return findKthLargestGO(arr, p, q-1, k)
	}

	//k > q+1, 说明第k大的数在右边区间, 也就是arr[q, r]
	//则个时候说明第k大的数在arr[q+1,r]区间里面
	return findKthLargestGO(arr, q+1, r, k)

}
