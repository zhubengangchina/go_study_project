package main

import (
	"errors"
	"fmt"
)

/**
题目：实现一个整数栈（Stack）
✨ 功能要求：
支持以下操作：

Push(int)：压栈

Pop() (int, error)：弹栈，空栈时返回错误

Peek() (int, error)：查看栈顶元素但不弹出

IsEmpty() bool：判断是否为空

可选加分项：

支持最大容量限制

实现 String() 方法，美化输出
*/

// Stack 是一个整数栈
type Stack struct {
	data     []int
	capacity int
}

// NewStack 创建一个新的栈，capacity 为 0 表示无限容量
func NewStack(capacity int) *Stack {
	return &Stack{
		data:     []int{},
		capacity: capacity,
	}
}

// Stack 结构体方法 push
func (s *Stack) Push(val int) error {
	if s.capacity > 0 && len(s.data) >= s.capacity {
		return errors.New("stack overflow: reach capacity limit")
	}
	s.data = append(s.data, val)
	return nil
}

//Pop() (int, error)：弹栈，空栈时返回错误

func (s *Stack) Pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack underflow: empty stack")
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, nil
}

// Peek() (int, error)：查看栈顶元素但不弹出
func (s *Stack) Peek() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("cannot peek: stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty() bool：判断是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

// 美化输出
func (s *Stack) String() string {
	return fmt.Sprintf("Stack: %v (top → bottom)", reverseCopy(s.data))
}

func reverseCopy(slices []int) []int {
	n := len(slices)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = slices[n-1-i]
	}
	return out

}

func main() {
	stack := NewStack(3)
	fmt.Println(stack.Push(10))
	fmt.Println(stack.Push(11))
	fmt.Println(stack.Push(12))
	fmt.Println(stack.Push(13))

	fmt.Println(stack.String())

	top, _ := stack.Peek()
	fmt.Println("Top:", top)

	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		fmt.Println("Top:", val)
	}

	_, err := stack.Pop()
	fmt.Println("Error:", err)

}
