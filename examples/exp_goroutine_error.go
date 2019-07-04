package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

type ApiData string
type apiError error

type DecoratorResult struct {
	ApiData  ApiData
	ApiError apiError
}

// 模拟外部请求（正常场景）
func doApiSuccess() (ApiData, apiError) {
	//time.Sleep(time.Second)
	return "api data", nil
}

// 模拟外部请求（异常场景）
func doApiFailure() (ApiData, apiError) {
	time.Sleep(time.Second)
	return "", errors.New("api failure")
}

// 模拟外部请求（超时场景）
func doApiTimeout() (ApiData, apiError) {
	time.Sleep(10 * time.Second)
	return "", errors.New("api timeout")
}

// 接口异步请求装饰器
func apiDecorator(apiFunc func() (ApiData, apiError), decoratorResultChan chan DecoratorResult, timeout int) {
	d := time.Duration(timeout) * time.Second
	timeAfter := time.After(d)

	fmt.Println("[request ] start")
	done := make(chan bool) // 外部接口请求完成

	go func(decoratorResultChan chan DecoratorResult) {
		res, err := apiFunc()
		decoratorResultChan <- DecoratorResult{res, err}
		done <- true
	}(decoratorResultChan)

	fmt.Println("[request ] end")

	select {
	case <-done:
		fmt.Println("[request ] get response!")
		return
	case <-timeAfter:
		fmt.Println("[request ] timeout...")
		decoratorResultChan <- DecoratorResult{"", errors.New("timeout")}
		return
	}
}

// 接口异步结果处理器
func apiResultHandler(decoratorResultChan chan DecoratorResult, done chan bool) {
	for {
		select {
		case decoratorResult := <-decoratorResultChan:
			fmt.Println("[response]", "data:", decoratorResult.ApiData, "; error:", decoratorResult.ApiError)
			done <- true
			return
		default:
			fmt.Println("[response] no message received")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 普通方式
	//_, e := doApiFailure()
	//if e != nil{
	//	fmt.Println(e.Erroror())
	//}

	timeoutRequest := 5
	resultChan := make(chan DecoratorResult, 1)

	// 同步协程方式
	//go func(decoratorResultChan chan DecoratorResult) {
	//	res, err := doApiFailure()
	//	decoratorResultChan <- DecoratorResult{res, err}
	//	close(decoratorResultChan)
	//}(resultChan)
	//
	//apiResult := <-resultChan
	//if apiResult.ApiError != nil {
	//	fmt.Println(apiResult.ApiError)
	//}

	// 异步协程方式
	forever := make(chan bool)

	// 外部请求 - 异步发起（三种场景：正常、失败、超时）
	//go apiDecorator(doApiSuccess, resultChan, timeoutRequest)
	go apiDecorator(doApiFailure, resultChan, timeoutRequest)
	//go apiDecorator(doApiTimeout, resultChan, timeoutRequest)

	// 外部请求 - 异步结果
	go apiResultHandler(resultChan, forever)

	fmt.Printf("[*] Waiting for logs. To exit press CTRL+C\n")

	<-forever
}
