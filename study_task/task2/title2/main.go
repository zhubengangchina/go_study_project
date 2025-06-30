package main

import (
	"fmt"
	"sync"
	"time"
)

/*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

type Task func()

// 调度器结构体
type Schedule struct {
	tasks []Task
}

// 添加任务
func (s *Schedule) AddTask(t Task) {
	s.tasks = append(s.tasks, t)
}

// 运行任务
func (s *Schedule) Run() {
	var wg sync.WaitGroup
	for i, task := range s.tasks {
		wg.Add(1)
		go func(id int, t Task) {
			defer wg.Done()
			start := time.Now()
			t()
			duration := time.Since(start)
			fmt.Printf("任务 %d 执行时间: %v\n", id+1, duration)
		}(i, task)
	}
	wg.Wait()
}

func main() {
	s := &Schedule{}
	s.AddTask(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("任务 1 完成")
	})

	s.AddTask(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务 2 完成")
	})

	s.AddTask(func() {
		time.Sleep(700 * time.Millisecond)
		fmt.Println("任务 3 完成")
	})

	s.Run()
}
