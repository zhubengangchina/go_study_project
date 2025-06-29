package main

import (
	"fmt"
	"sync"
)

/*
并发安全的缓存（Map + Mutex）
🎯 题目目标：
实现一个线程安全的内存缓存系统，支持并发读写。

✅ 功能要求
提供以下接口：

Set(key string, value string) 设置缓存值

Get(key string) (string, bool) 获取缓存值（返回值 + 是否存在）

Delete(key string) 删除键

支持多个 goroutine 并发读写缓存而不会发生数据竞争

使用 sync.Mutex 或 sync.RWMutex 实现并发安全

*/

type SafeCache struct {
	mu    sync.RWMutex
	store map[string]string
}

// set 插入或更新缓存
func (s *SafeCache) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

// get 查询缓存
func (s *SafeCache) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	val, ok := s.store[key]
	return val, ok
}

// delete 删除缓存
func (s *SafeCache) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

func main() {

	sc := &SafeCache{
		store: make(map[string]string),
	}

	var wg sync.WaitGroup
	//并发写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			sc.Set(key, value)
			fmt.Printf("Set %s = %s\n", key, value)
		}(i)
	}

	//并发读取
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", i)
			if value, ok := sc.Get(key); ok {
				fmt.Printf("Get %s = %s\n", key, value)
			} else {
				fmt.Printf("Get %s = not found\n", key)
			}
		}(i)
	}

	wg.Wait()

}
