package main

import "fmt"

/*
题目描述
编写一个函数，查找字符串数组中的最长公共前缀

如果不存在公共前缀，返回空字符串 ""。
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, val := range strs[1:] {
		i := 0
		//比较prefix和当前 val
		for i < len(prefix) && i < len(val) && prefix[i] == val[i] {
			i++
		}
		prefix = prefix[:i] //缩短prefix
		if prefix == "" {
			return ""
		}

	}
	return prefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))          // "fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))             // ""
	fmt.Println(longestCommonPrefix([]string{"interview", "internet", "internal"})) // "inte"
	fmt.Println(longestCommonPrefix([]string{}))                                    // ""
}
