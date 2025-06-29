package main

import (
	"fmt"
	"sync"
)

/*
题目：并发安全的计数器
🧩 功能要求：
实现一个结构体 Counter，提供如下方法：

Inc()：增加计数

Dec()：减少计数

Get() int：获取当前值

必须支持 并发安全，也就是说：

多个 goroutine 同时操作，不会出现竞态或数据错误

限制条件：

不允许使用全局变量

所有数据应封装在结构体中

*/

type Counter struct {
	mu  sync.Mutex
	val int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) Dec() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val--

}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

func main() {
	counter := &Counter{}
	var wg sync.WaitGroup

	// 启动 100 个 goroutine 同时递增
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()

	}

	// 再启动 50 个 goroutine 同时递减
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			counter.Dec()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("最终计算值：", counter.Get())

}
