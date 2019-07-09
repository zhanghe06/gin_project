# Golang

https://golang.org


### install （Mac OS）

https://golang.org/dl/

```
wget https://dl.google.com/go/go1.12.6.darwin-amd64.pkg
```

安装

https://golang.org/doc/install

设置环境变量 .zshrc
```
# golang
export GOROOT=/usr/local/go
export GOPATH=$HOME/work
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```


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


### json

Encode  
将一个对象编码成JSON数据，接受一个interface{}对象，返回[]byte和error：  
func Marshal(v interface{}) ([]byte, error)  

Marshal函数将会递归遍历整个对象，依次按成员类型对这个对象进行编码，类型转换规则如下：  
```
bool类型 转换为JSON的Boolean
整数，浮点数等数值类型 转换为JSON的Number
string 转换为JSON的字符串(带""引号)
struct 转换为JSON的Object，再根据各个成员的类型递归打包
数组或切片 转换为JSON的Array
[]byte 会先进行base64编码然后转换为JSON字符串
map 转换为JSON的Object，key必须是string
interface{} 按照内部的实际类型进行转换
nil 转为JSON的null
channel,func等类型 会返回UnsupportedTypeError
```

interface{}类型在Unmarshal时，会自动将JSON转换为对应的数据类型：
```
JSON的boolean 转换为bool
JSON的数值 转换为float64
JSON的字符串 转换为string
JSON的Array 转换为[]interface{}
JSON的Object 转换为map[string]interface{}
JSON的null 转换为nil
```



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

使用场景（把channel用在数据流动的地方）
- 消息传递、消息过滤
- 信号广播
- 事件订阅与广播
- 请求、响应转发
- 任务分发
- 结果汇总
- 并发控制
- 同步与异步
- ...


1、相对sync.WaitGroup而言，golang中利用channel实习同步则简单的多．channel自身可以实现阻塞，其通过<-进行数据传递，channel是golang中一种内置基本类型，对于channel操作只有４种方式：

- 创建channel(通过make()函数实现，包括无缓存channel和有缓存channel);
- 向channel中添加数据（channel<-data）;
- 从channel中读取数据（data<-channel）;
- 关闭channel(通过close()函数实现，关闭之后无法再向channel中存数据，但是可以继续从channel中读取数据）

2、channel分为有缓冲channel和无缓冲channel,两种channel的创建方法如下:

- var ch = make(chan int) //无缓冲channel,等同于make(chan int ,0)
- var ch = make(chan int,10) //有缓冲channel,缓冲大小是10

需要注意 go协程与主协程的顺序，防止死锁

3、close主要用来关闭channel通道其用法为close(channel)，并且是在生产者的地方关闭channel，而不是在消费者的地方关闭．并且关闭channel后，便不可再向channel中继续存入数据，但是可以继续从channel中读取数据

需要显示关闭channel的场景：生产者、消费者对channel数量无法提前定义，只能在生产完成时，由生产者关闭channel，消费者select接收关闭信号
```go
// stopCh并不需要传递任何数据
// 只是要给所有协程发送退出的信号
type Handler struct {
    stopCh chan struct{}
    reqCh chan *Request
}

func (h *Handler) Stop() {
    close(h.stopCh)

    // 可以使用WaitGroup等待所有协程退出
}

// 收到停止后，不再处理请求
func (h *Handler) loop() error {
    for {
        select {
        case req := <-h.reqCh:
            go handle(req)
        case <-h.stopCh:
            return
        }
    }
}
```
示例中，h.stopCh为空，case <-h.stopCh一直阻塞；一旦关闭h.stopCh，则case <-h.stopCh接收到nil

4、channel阻塞超时处理，通过select设置超时处理

读写特性
```
//定义只读的channel
read_only := make (<-chan int)

//定义只写的channel
write_only := make (chan<- int)

//可同时读写
read_write := make (chan int)
```

channel可进行3种操作：读、写和关闭；把这3种操作和3种channel状态可以组合出9种情况:

操作 | nil的channel | 正常channel | 已关闭channel
--- | --- | --- | ---
<- ch | 阻塞 | 成功或阻塞 | 读到零值
ch <- | 阻塞 | 成功或阻塞 | panic
close(ch) | panic | 成功 | panic

有缓冲通道可供多个协程同时处理，在一定程度可提高并发性。

若子协程读 channel，则主协程阻塞写

若子协程写 channel，则主协程阻塞读

语法上阻塞操作只能写在后面，不然子协程异步操作没法执行，也就是说不满足 chennel 通信条件


### select case

如果有同时多个case去处理,比如同时有多个channel可以接收数据，那么Go会伪随机的选择一个case处理(pseudo-random)。  
如果没有case需要处理，则会选择default去处理，如果default case存在的情况下。  
如果没有default case，则select语句会阻塞，直到某个case需要处理

select语句和switch语句一样，它不是循环，它只会选择一个case来处理，如果想一直处理channel，你可以在外面加一个无限的for循环


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


### fatal error & panic

- fatal error
    - 立即退出应用程序，defer函数不执行
- panic
    - 函数本身停止执行，defer函数被执行，注意不是应用程序停止
    - 需要recover()来辅助处理：丢弃或者上抛


### 新建

- new() 返回指针
- struct{} 返回struct
- make() 仅适用于 map，slice 和 channel
