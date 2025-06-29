package main

import "fmt"

/*
给定一个只包含 ()、[]、{} 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。

左括号必须以正确的顺序闭合。
*/
func isValid(s string) bool {

	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		//如果是又括号，检查匹配
		if v, ok := mapping[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1] //弹栈
		} else {
			//否则是左括号，压入栈
			stack = append(stack, ch)
		}

	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))           // true
	fmt.Println(isValid("()[]{}"))       // true
	fmt.Println(isValid("(]"))           // false
	fmt.Println(isValid("([)]"))         // false
	fmt.Println(isValid("{[]}"))         // true
	fmt.Println(isValid("(({{[[]]}}))")) // true
}
