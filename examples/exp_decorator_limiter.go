package main

import (
	"github.com/lunny/log"
	"time"
)

// 模拟接口请求发起
func doLimiterApiRequest() error {
	time.Sleep(time.Millisecond * 1000)
	log.Info("api request")
	return nil
}

/*
 * 并发限速装饰器
 */
func limiterApiDecorator(apiFunc func() error, concurrentMaxNum int) {
	limiter := make(chan int, concurrentMaxNum)
	for {
		limiter <- 1
		go func() {
			_ = apiFunc()
			<-limiter
		}()
	}
}

func main() {
	limiterApiDecorator(doLimiterApiRequest, 3)
}
