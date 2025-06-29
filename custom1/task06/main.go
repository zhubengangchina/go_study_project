package main

import (
	"context"
	"fmt"
	"time"
)

/*
使用 context 控制并发任务的取消与超时
🎯 题目目标：
你将实现一个支持取消和超时的任务系统，理解 Go 的 context 包如何控制 goroutine 生命周期。

✅ 功能要求：
定义一个 TaskRunner 结构，运行一个或多个“长时间运行”的任务（模拟：sleep + 打印）

支持：

✅ 主动取消任务（context.WithCancel）

✅ 设置任务超时（context.WithTimeout）

所有任务接收到取消信号后应立刻终止（模拟中断）

打印任务是否完成、是否被取消、是否超时
*/

type TaskRunner struct{}

func NewTaskRunner() *TaskRunner {
	return &TaskRunner{}
}

// run 启动一个任务，支持context控制
func (tx *TaskRunner) Run(ctx context.Context, name string, duration time.Duration) {
	fmt.Printf("%s started \n", name)
	select {
	case <-time.After(duration):
		fmt.Printf("%s completed\n", name)
	case <-ctx.Done():
		fmt.Printf("%s canceled or timeout: %v\n", name, ctx.Err())
	}
}

func main() {
	taskRunner := NewTaskRunner()
	// 示例：任务持续 5 秒，但 2 秒后超时取消
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	taskRunner.Run(ctx, "Task-A", 5*time.Second)
}
