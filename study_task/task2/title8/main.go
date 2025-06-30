package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func main() {
	var counter int64 //使用 int64 类型以配合 atomic 包
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子递增
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终计数器的值:", counter)
}
