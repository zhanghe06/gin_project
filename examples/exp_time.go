package main

import (
	"fmt"
	"time"
)

/*
2019-07-03 14:01:12.371795 +0800 CST m=+0.000158910
1562133672
2019-07-03 14:01:12
*/
func main() {
	now := time.Now()
	fmt.Println(now)

	nowUnix := time.Now().Unix()
	fmt.Println(nowUnix)

	// 当前格式化时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5
}
