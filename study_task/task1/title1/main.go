package main

import "fmt"

//136. 只出现一次的数字

func singleNumber(nums []int) int {
	counterMap := make(map[int]int)

	//统计每个数字出现的次数
	for _, v := range nums {
		counterMap[v]++
	}

	for num, count := range counterMap {
		if count == 1 {
			return num
		}
	}
	return -1
}
func main() {

	nums := []int{2, 3, 2, 4, 3}
	result := singleNumber(nums)
	fmt.Println("只出现一次的数字:", result)
}
