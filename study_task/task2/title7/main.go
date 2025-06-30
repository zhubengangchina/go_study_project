package main

import (
	"fmt"
	"sync"
)

/*
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。

有一个共享变量 counter；

启动 10 个协程，每个协程递增 counter 1000 次；

使用 sync.Mutex 确保并发安全；

最终输出 counter 的正确值：10000。
*/

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("最终计数器的值:", counter)
}
