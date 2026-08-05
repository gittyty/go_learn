package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(1 * time.Millisecond) // 模拟轻微耗时
	}
}

func main() {
	// 设置 P 的数量为 2（强制限制并行度，方便观察调度轨迹）
	runtime.GOMAXPROCS(2)

	// 启动 4 个协程
	go say("Hello")
	go say("World")
	go say("Go!")
	go say("调度中...")

	// 主协程睡 1 秒，等待其他协程跑完
	time.Sleep(1 * time.Second)
	fmt.Println("主程序结束")
}
