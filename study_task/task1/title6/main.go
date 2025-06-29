package main

import "fmt"

/*
题目描述：
给定一个有序数组 nums，原地删除重复出现的元素，使每个元素只出现一次，并返回新的长度。

不允许使用额外数组空间，空间复杂度必须为 O(1)。
解题思路：双指针法
i 是慢指针，指向最后一个不重复元素的位置

j 是快指针，负责遍历整个数组

当 nums[i] != nums[j]，就说明 nums[j] 是一个新的不重复元素，将它移动到 i+1 位置，然后 i++
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	length := removeDuplicates(nums)
	fmt.Println("新长度:", length)
	fmt.Println("去重后的数组:", nums[:length]) // 输出有效部分
}
