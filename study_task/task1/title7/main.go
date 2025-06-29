package main

import (
	"fmt"
	"sort"
)

/*
给你一个区间数组 intervals，每个区间是形如 [start, end] 的二维数组，合并所有重叠的区间，返回一个不重叠的区间数组
解题思路
排序： 按照每个区间的起始位置 start 进行升序排序；

遍历合并：

如果当前区间的 start ≤ 上一个区间的 end，说明重叠，合并；

否则，直接加入结果集。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	//按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//初始化结果切片
	merged := [][]int{intervals[0]}

	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			// 重叠，合并
			if curr[1] >= last[1] {
				last[1] = curr[1]
			}
		} else {
			//不重叠  直接追加
			merged = append(merged, curr)
		}
	}
	return merged
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	result := merge(intervals)
	fmt.Println("合并后的区间：", result)
}
