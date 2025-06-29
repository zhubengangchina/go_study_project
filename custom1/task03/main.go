package main

import (
	"errors"
	"fmt"
	"time"
)

/*
题目：带缓冲的消息队列
🧩 功能要求：
使用 Go 的 chan 实现一个简单的消息队列，支持以下操作：

Send(msg string) error：发送消息

Receive() (string, error)：接收消息

内部使用带缓冲的 chan string（例如容量为 10）

阻塞行为由 channel 控制，不需要自己实现等待机制

✅ 加分项：
添加超时机制：当发送或接收阻塞超过 1 秒时返回超时错误

添加 Close() 方法，并支持接收方检测已关闭的状态
*/

type MessageQueue struct {
	ch    chan string
	close chan struct{} //关闭标志
}

// 创建一个固定容量的消息队列
func NewMessageQueue(capacity int) *MessageQueue {
	return &MessageQueue{
		ch:    make(chan string, capacity),
		close: make(chan struct{}),
	}
}

func (m *MessageQueue) Send(msg string) error {
	select {
	case m.ch <- msg:
		return nil
	case <-time.After(1 * time.Second):
		return errors.New("send timeout")
	case <-m.close:
		return errors.New("send failed: queue closed")
	}
}

func (m *MessageQueue) Receive() (string, error) {
	select {
	case msg := <-m.ch:
		return msg, nil
	case <-time.After(1 * time.Second):
		return "", errors.New("Receive timeout")
	case <-m.close:
		return "", errors.New("Receive failed: queue closed")
	}

}

func (m *MessageQueue) Close() {

	select {
	case <-m.close:
	default:
		close(m.close)
	}

}

func main() {
	mq := NewMessageQueue(3)
	// 启动一个发送者
	go func() {
		for i := 0; i < 5; i++ {
			err := mq.Send(fmt.Sprintf("msg: %d", i))
			if err != nil {
				fmt.Println("Send error:", err)
			}
			time.Sleep(200 * time.Millisecond)
		}
	}()

	// 启动一个接收者
	go func() {
		for {
			msg, err := mq.Receive()
			if err != nil {
				fmt.Println("Receive error:", err)
				return
			}
			fmt.Println("Received:", msg)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	//稍后关闭队列,,避免主线程结束直接挂壁子协程
	time.Sleep(3 * time.Second)
	mq.Close()
	time.Sleep(1 * time.Second)
}
