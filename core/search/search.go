package search

import (
	"fmt"
	"math"
)

// CompareF2 比较函数, 返回0表示两个元素相同, 负数表示a < b, 正数表示 a > b
type CompareF2 func(a, b interface{}) int

// BinarySearch 二分查找
// 时间复杂度O(logn)
// @return 返回匹配元素的下标, 无匹配时返回-1
func BinarySearch(arr []interface{}, target interface{}, cmp CompareF2) int {
	p := 0
	r := len(arr) - 1
	index := -1
	for p <= r {
		fmt.Println(arr[p : r+1])
		//先计算二分查找的中间数
		median := p + (r-p)/2
		if cmp(arr[median], target) == 0 {
			index = median
			break
		}
		if cmp(target, arr[median]) < 0 {
			r = median - 1
		} else if cmp(target, arr[median]) > 0 {
			p = median + 1
		}
	}
	return index
}

// Square 求n的平方根
// 采用二分查找思想解决问题
func Square(n float64) float64 {
	//从0-n开始, 选取中间数 m, 比较m^2和n, 如果等于n, 说明m就是目标数, 如果小于n, 则说明目标在0-m区间, 否则说明目标在m-n区间
	p := 0.0
	r := n
	m := p + (r-p)/2
	for math.Abs(m*m-n) > 0.000001 {
		fmt.Printf("%f * %f = %f\n", m, m, m*m)
		if m*m < n {
			p = m
		} else if m*m > n {
			r = m
		}
		m = p + (r-p)/2
	}

	return m
}

// BinarySearch1 查找第一个等于给定值的元素
// 时间复杂度O(logn)
// @return 返回匹配元素的下标, 无匹配时返回-1
func BinarySearch1(arr []interface{}, target interface{}, cmp CompareF2) int {
	low := 0
	high := len(arr) - 1
	index := -1
	for low <= high {
		fmt.Println(arr[low : high+1])
		//先计算二分查找的中间数
		median := low + (high-low)/2
		if cmp(target, arr[median]) < 0 {
			high = median - 1
		} else if cmp(target, arr[median]) > 0 {
			low = median + 1
		} else {
			if median == 0 || cmp(target, arr[median-1]) != 0 { //如果median等于0, 说明arr[median]是数组第一个元素, 肯定是第一个符合条件的元素. 或者arr[median-1]不符合条件,则说明找到第一个符合条件的元素了. 以上两种情况都不是,则说明区间目标应该是在区间[low,median-1],循环继续.
				index = median
				break
			} else {
				high = median - 1
			}
		}
	}
	return index
}

// BinarySearch2 查找最后一个值等于给定值的元素
// 时间复杂度O(logn)
// @return 返回匹配元素的下标, 无匹配时返回-1
func BinarySearch2(arr []interface{}, target interface{}, cmp CompareF2) int {
	low := 0
	high := len(arr) - 1
	index := -1
	for low <= high {
		fmt.Println(arr[low : high+1])
		//先计算二分查找的中间数
		median := low + (high-low)/2
		if cmp(target, arr[median]) < 0 {
			high = median - 1
		} else if cmp(target, arr[median]) > 0 {
			low = median + 1
		} else {
			if median == len(arr)-1 || cmp(target, arr[median+1]) != 0 { //如果arr[median]是数组最后一个元素, 那肯定是最后一个等于目标的元素. 或者arr[median+1]不符合条件,则说明找到最后一个等于目标的元素了. 以上两种情况都不是,则说明区间目标应该是在区间[median+1,high],循环继续.
				index = median
				break
			} else {
				low = median + 1
			}
		}
	}
	return index
}

// BinarySearch3 查找第一个大于等于给定值的元素
// 时间复杂度O(logn)
// @return 返回匹配元素的下标, 无匹配时返回-1
func BinarySearch3(arr []interface{}, target interface{}, cmp CompareF2) int {
	low := 0
	high := len(arr) - 1
	index := -1
	for low <= high {
		fmt.Println(arr[low : high+1])
		//先计算二分查找的中间数
		median := low + (high-low)/2
		if cmp(arr[median], target) >= 0 {
			if median == 0 || cmp(arr[median-1], target) < 0 { //如果median等于0, 说明arr[median]是数组第一个元素, 肯定就是第一个大于等于目标的元素. 或者arr[median-1]小于目标,则说明找到第一个符合条件的元素了. 以上两种情况都不是,则说明区间目标应该是在区间[low,median-1],循环继续.
				index = median
				break
			} else {
				high = median - 1
			}
		} else {
			low = median + 1
		}
	}
	return index
}

// BinarySearch4 查找最后一个小于等于给定值的元素
// 时间复杂度O(logn)
// @return 返回匹配元素的下标, 无匹配时返回-1
func BinarySearch4(arr []interface{}, target interface{}, cmp CompareF2) int {
	low := 0
	high := len(arr) - 1
	index := -1
	for low <= high {
		fmt.Println(arr[low : high+1])
		//先计算二分查找的中间数
		median := low + (high-low)/2
		if cmp(arr[median], target) <= 0 {
			if median == len(arr)-1 || cmp(arr[median+1], target) > 0 { //如果arr[median]是数组最后一个元素, 那肯定就是最后一个小于等于目标的元素. 或者arr[median+1]大于目标,则说明找到最后一个符合条件的元素了. 以上两种情况都不是,则说明区间目标应该是在区间[median+1,high],循环继续.
				index = median
				break
			} else {
				low = median + 1
			}
		} else {
			high = median - 1
		}
	}

	return index
}

// LetCodeBinarySearch LetCode第33题, 在循环数组中作二分查找, 复杂度O(logn)
// nums元素不重复, 而且是循环有序, 比如45670123
func LetCodeBinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	index := -1
	for low <= high {
		fmt.Println(nums[low : high+1])
		//先计算二分查找的中间数
		median := low + (high-low)/2
		if nums[median] == target {
			index = median
			break
		}

		// 当median左边区间第一个元素low小于median, 假设左边区间存在一个大于median的数, 那么数组即存在[小小..大..小(median)..]这样的结构, 此时不可能是循环数组了.
		// 因此当low小于median时, 说明median左边的元素都小于median
		//
		// 当low等于median时, 则说明数组只剩下2个元素了.[大, 小] 或者 [小, 大].
		if nums[low] <= nums[median] {
			// 这里注意, 虽然median左边的元素都比median小, 但是不一定都比target小. 比如 4567012,寻找2时,456都比7小,但是却比2大. 因此需要确保target比median小的同时比low大
			if nums[low] <= target && target <= nums[median] {
				high = median - 1
			} else {
				low = median + 1
			}
		} else { // 同理, 当low大于median时, 假设右边区间存在一个小于median的数, 则数组存在[大..小(median)..小小..]这样的结构, 此时也不可能是循环数组了. 因此当low大于median时, 说明右边区间的元素都大于median

			// 这里同样需要注意, 虽然右边区间的元素比median大, 但是不一定都比target大. 比如6701234,寻找7时, 234都比1大, 但是却比7小. 因此需要确保target比median大的同时比high小
			if nums[median] <= target && target <= nums[high] {
				low = median + 1
			} else {
				high = median - 1
			}
		}
	}

	return index
}
