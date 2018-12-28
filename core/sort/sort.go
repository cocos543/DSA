package sort

import (
	"fmt"
)

// Swap 交换in中下标i和j对应元素的位置
// func Swap(in interface{}, i, j int) {
// 	v := reflect.ValueOf(in)
// 	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
// 		fmt.Println(v.Kind())
// 		panic("Swap input's type must be a Slice or Array")
// 	}

// 	temp := v.Index(j).Interface()
// 	v.Index(j).Set(v.Index(i))
// 	v.Index(i).Set(reflect.ValueOf(temp))
// }

// Swap 交换in中下标i和j对应元素的位置
func Swap(in []interface{}, i, j int) {
	temp := in[i]
	in[i] = in[j]
	in[j] = temp
}

// BubbleSort 冒泡排序
// 时间复杂度O(n^2), 空间复杂度O(1).
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func BubbleSort(arr []interface{}, compare func(a, b interface{}) bool) {
	len := len(arr)
	for i := 0; i < len; i++ {
		skip := true
		for j := 0; j < len-i-1; j++ {
			if !compare(arr[j], arr[j+1]) {
				Swap(arr, j, j+1)
				//没有交换的时候, 说明已经排完了
				skip = false
			}
		}
		if skip {
			return
		}
		fmt.Println(arr)
	}
}

// StraightInsertionSort 直接插入排序
// 时间复杂度O(n^2), 空间复杂度O(1).
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func StraightInsertionSort(arr []interface{}, compare func(a, b interface{}) bool) {
	fmt.Println("init:", arr)
	len := len(arr)
	//i表示未排序区间的起始位置
	for i := 1; i < len; i++ {
		//[0,i-1]为已排序区间
		j := i - 1
		for ; j >= 0; j-- {
			if !compare(arr[j], arr[j+1]) {
				//arr[j]就是arr[i]需要插入的位置
				//需要腾出位置出来, 让arr[i]插入
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				fmt.Println(arr)
			} else {
				//因为左边区间都是有序的, 所以当发现j+1不需要和j交换位置时, 那么说明j+1也不需要和[0, j)之间任何一个元素交换
				//j+1这个元素的当前位置就是最终位置, 可以跳出循环
				break
			}
		}
	}
}
