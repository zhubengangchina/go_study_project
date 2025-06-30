package main

import (
	"fmt"
	"time"
)

/*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。

使用一个协程生产整数（1~10）并发送到通道中；

使用另一个协程从通道中接收整数并打印；

使用 close(channel) 正确关闭通道；

使用 range 读取通道直到关闭。
*/

func sendNumber(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	//发送完关闭管道
	close(ch)
}

func receiveNumber(ch chan int) {
	for val := range ch {
		fmt.Println("接收到数字:", val)
	}
}

func main() {
	ch := make(chan int) //无缓冲管道

	go sendNumber(ch)
	go receiveNumber(ch)

	//保证协程全部执行完 主线程才结束
	time.Sleep(2 * time.Second)

}
