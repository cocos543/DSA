package sort

import "fmt"

// MergingSort 归并排序
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func MergingSort(arr []interface{}, compare CompareF) {
	mergingSortGo(arr, 0, len(arr)-1, compare)

}

// mergingSortGo 归并排序的真实算法部分. p为数组的起始元素下标, r为数组的末尾元素下标
// 时间复杂度O(nlogn), 空间复杂度O(n)
func mergingSortGo(arr []interface{}, p, r int, compare CompareF) {
	if p >= r {
		return
	}

	q := p + (r-p)/2
	mergingSortGo(arr, p, q, compare)
	mergingSortGo(arr, q+1, r, compare)
	merge(arr, p, q, r, compare)

}

// merge 将 arr[p, q] 与 arr[q+1, r]进行合并为有序数组
func merge(arr []interface{}, p, q, r int, compare CompareF) {
	fmt.Printf("arr[%v,%v] arr[%v,%v]\n", p, q, q+1, r)
	// 申请一个临时数组, 用于存放已经排序的元素
	temp := make([]interface{}, 0)
	// i和j分别指向两个数组的起始位置
	i, j := p, q+1

	for i <= q && j <= r {
		// 这里两个数相同时, compare函数返回true就表示不需要交换位置了. 也就是稳定算法
		if compare(arr[i], arr[j]) {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	// 把还有剩余元素的数组直接追加到临时数组末尾, 因为剩余元素也是有序的, 所以最终临时数组是有序的
	if i <= q {
		temp = append(temp, arr[i:q+1]...)
	} else {
		temp = append(temp, arr[j:r+1]...)
	}

	for i := p; i <= r; i++ {
		arr[i] = temp[i-p]
	}
	fmt.Printf("%v\n\n", arr)
}
