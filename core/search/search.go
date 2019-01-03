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
