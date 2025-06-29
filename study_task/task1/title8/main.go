package main

import "fmt"

/*
两数之和

考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int) //值 =》下标

	for index, num := range nums {
		result := target - num
		if val, ok := seen[result]; ok {
			return []int{val, index}
		}
		seen[num] = index
	}

	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("结果下标:", result) // [0, 1]
}
