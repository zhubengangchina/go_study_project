package main

import (
	"fmt"
	"sync"
)

/*
题目目标：
实现一个工作池（worker pool），能并发处理任务，并控制最大并发数。

✅ 功能要求：
有一组任务（如计算整数的平方）

使用固定数量的 worker 来并发处理这些任务

主线程等待所有任务完成后退出（使用 sync.WaitGroup）

📦 示例设定：
我们有 100 个数字，任务是对每个数字求平方，由最多 5 个 worker 并发处理。
*/

// 处理函数
func processTask(n int) int {
	return n * n
}

// 工作协程
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range jobs {
		fmt.Printf("Worker %d processing %d\n", id, n)
		val := processTask(n)
		results <- val
	}
}

func main() {
	numWorks := 5
	numJobs := 20

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	//启动固定数量的worker 来处理任务
	for i := 1; i <= numWorks; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)

	}

	//发送任务
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	//等待所有worker完成
	go func() {
		wg.Wait()
		close(results)
	}()

	//主线程收集结果
	for result := range results {
		fmt.Println(result)
	}
}
