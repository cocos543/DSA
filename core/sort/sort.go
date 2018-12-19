package sort

import (
	"fmt"
	"reflect"
)

// Swap 交换in中下标i和j对应元素的位置
func Swap(in interface{}, i, j int) {
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		fmt.Println(v.Kind())
		panic("Swap input's type must be a Slice or Array")
	}

	temp := v.Index(j).Interface()
	v.Index(j).Set(v.Index(i))
	v.Index(i).Set(reflect.ValueOf(temp))
}

// BubbleSort 冒泡排序
// 时间复杂度O(n^2), 空间复杂度O(1).
// compare函数, 返回true表示不需要交换位置, false表示需要交换位置.
func BubbleSort(arr []int, compare func(a, b interface{}) bool) {
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
