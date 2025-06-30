package main

import (
	"fmt"
	"sync"
)

/*
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/

func sendNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i //如果通道满了，阻塞直到有空位
	}
}

func receiveNumber(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("接收到数字:", val)
	}
}

func main() {
	ch := make(chan int, 20) //带20 缓冲的管道

	var wg sync.WaitGroup
	wg.Add(2)

	go sendNumber(ch, &wg)
	go receiveNumber(ch, &wg)

	wg.Wait()
}
