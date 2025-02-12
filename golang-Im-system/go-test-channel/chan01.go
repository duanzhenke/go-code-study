package go_test_channel

import (
	"fmt"
)

func sendData(ch chan int) {
	// 模拟一些耗时操作
	fmt.Println("Sending data...")
	ch <- 42 // 发送数据到通道
}

func main() {
	ch := make(chan int) // 创建一个无缓冲通道

	go sendData(ch) // 启动一个新的 goroutine 来发送数据

	// 主 goroutine 阻塞在这里，直到从通道接收到数据
	data := <-ch
	fmt.Println("Received data:", data)
}
