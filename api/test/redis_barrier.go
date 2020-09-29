package main

import (
	"fmt"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/core/syncx"
	"time"
)

func main() {
	const round = 5
	barrier := syncx.NewSharedCalls() //进程内共享调用,注意不是协程内，因为是多个协程同一时间内，拿到相同的返回

	for i := 0; i < round; i++ {
		// 多个线程同时执行
		go func() {
			// 可以看到，多个线程在同一个key上去请求资源，获取资源的实际函数只会被调用一次
			val, err := barrier.Do("once", func() (interface{}, error) {
				// sleep 1秒，为了让多个线程同时取once这个key上的数据
				time.Sleep(time.Second)
				// 生成了一个随机的id
				return stringx.RandId(), nil
			})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}
	ch := make(chan bool)
	<-ch
	for {
		time.Sleep(time.Second)
	}
}
