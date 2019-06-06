# Golang

https://golang.org


### strings

```go
import "strings"

startsWith := strings.HasPrefix("prefix", "pre") // true
endsWith := strings.HasSuffix("suffix", "fix") // true
```


### slice

range 下标从0开始计数


### map

map 非线程安全，不能并发写
```
fatal error: concurrent map writes
```

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) New() {
	ua.ages = make(map[string]int)
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

func main()  {
	var ua UserAges
	ua.New()
	go ua.Add("a", 1)
	go ua.Add("a", 2)
	go ua.Add("a", 3)
	go ua.Add("a", 4)
	go ua.Add("a", 5)
	time.Sleep(time.Second)
	go fmt.Println(ua.Get("a"))
	go fmt.Println(ua.Get("a"))
	go fmt.Println(ua.Get("a"))
	fmt.Println(ua.Get("b"))
	time.Sleep(time.Second)
}
```

### interface

interface{}变量的本质。
在Go语言中，一个interface{}类型的变量包含了2个指针，一个指针指向值的类型，另外一个指针指向实际的值。

对于一个interface{}类型的nil变量来说，它的两个指针都是0。这是符合Go语言对nil的标准定义的。

避免将一个有可能为nil的具体类型的值赋值给interface变量


### struct


### 互斥锁（sync.Mutex）和读写互斥锁（sync.RWMutex）

```go
var (
    // 逻辑中使用的某个变量
    count int
    // 与变量对应的使用互斥锁
    countGuard sync.RWMutex
)
func GetCount() int {
    // 锁定
    countGuard.RLock()
    // 在函数退出时解除锁定
    defer countGuard.RUnlock()
    return count
}
```

并发调用`countGuard.RLock()`时，并不会发生阻塞


### 协程同步 sync.waitGroup

golang中有2种方式同步程序，一种使用channel，另一种使用锁机制

WaitGroup总共有三个方法：
- Add(delta int)
- Done()
- Wait()

```
Add:添加或者减少等待goroutine的数量
Done:相当于Add(-1)
Wait:执行阻塞，直到所有的WaitGroup数量变成0
```

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    for i := 0; i > 5; i = i + 1 {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
            // defer wg.Add(-1)
            EchoNumber(n)
        }(i)
    }

    wg.Wait()
}

func EchoNumber(i int) {
    time.Sleep(3e9)
    fmt.Println(i)
}
```

```go
package main

import (
    "fmt"
    "sync"
)

//声明一个全局变量
var waitGroup sync.WaitGroup

func Afunction(shownum int) {
    fmt.Println(shownum)
    waitGroup.Done() // 任务完成，将任务队列中的任务数量-1，其实.Done就是.Add(-1)
}

func main() {
    for i := 0; i < 10; i++ {
        waitGroup.Add(1) // 每创建一个goroutine，就把任务队列中任务的数量+1
        go Afunction(i)
    }
    waitGroup.Wait() // .Wait()这里会发生阻塞，直到队列中所有的任务结束就会解除阻塞
}
```


### channel

1、相对sync.WaitGroup而言，golang中利用channel实习同步则简单的多．channel自身可以实现阻塞，其通过<-进行数据传递，channel是golang中一种内置基本类型，对于channel操作只有４种方式：

- 创建channel(通过make()函数实现，包括无缓存channel和有缓存channel);
- 向channel中添加数据（channel<-data）;
- 从channel中读取数据（data<-channel）;
- 关闭channel(通过close()函数实现，关闭之后无法再向channel中存数据，但是可以继续从channel中读取数据）

2、channel分为有缓冲channel和无缓冲channel,两种channel的创建方法如下:

- var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0)
- var ch = make(chan int,10) //有缓冲channel,缓冲大小是10

需要注意 go协程与主协程的顺序，防止死锁

3、close主要用来关闭channel通道其用法为close(channel)，并且实在生产者的地方关闭channel，而不是在消费者的地方关闭．并且关闭channel后，便不可再想channel中继续存入数据，但是可以继续从channel中读取数据

4、channel阻塞超时处理，通过select设置超时处理


### context

```
go get golang.org/x/net/context
```


### signal

有两种信号不能被拦截和处理: `SIGKILL`和`SIGSTOP`

当接收到信号时，进程会根据信号的响应动作执行相应的操作，信号的响应动作有以下几种：

- 中止进程(Term)
- 忽略信号(Ign)
- 中止进程并保存内存信息(Core)
- 停止进程(Stop)
- 继续运行进程(Cont)
