package main

import "fmt"

//回文数
//判断一个整数是否是回文数。正序（从左到右）和倒序（从右到左）读都是一样的整数

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0
	for x != 0 {
		fmt.Println(x, "-----", reversed)
		digit := x % 10
		reversed = reversed*10 + digit
		fmt.Println(x, "-----", reversed)
		x /= 10
	}
	return original == reversed
}

func main() {
	ok := isPalindrome(121)
	fmt.Println("是否回文数:", ok)
}
