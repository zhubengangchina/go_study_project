package main

import "fmt"

/*
加一

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i > 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits //无需进位，直接返回
		}
		digits[i] = 0
	}
	// 如果全是9，例如 [9,9,9]，结果应是 [1,0,0,0]
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))    // [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	fmt.Println(plusOne([]int{9, 9, 9}))    // [1 0 0 0]
	fmt.Println(plusOne([]int{0}))          // [1]
}
