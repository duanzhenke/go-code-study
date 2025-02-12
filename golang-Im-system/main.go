package main

import (
	"fmt"
)

func sendData(ch chan int) {
	// 模拟一些耗时操作
	fmt.Println("Sending data...")
	ch <- 42 // 发送数据到通道
}

func main() {
	ch := make(chan int, 1) // 创建一个有缓冲通道，缓冲区大小为1

	go sendData(ch) // 启动一个新的 goroutine 来发送数据

	// 主 goroutine 不会立即阻塞，因为通道有一个空闲的缓冲区位置
	data := <-ch
	fmt.Println("Received data:", data)
}
