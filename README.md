# gomq
消息队列工具-go实现


## Installation


```sh
$ go get github.com/wzcssw/gomq
```

## Quickstart

```go
    var queue = goqueue.NewInstance(goqueue.RedisConfig{Host: "localhost", Port: 6379, DB: 0})
    const key = "charactor" // 则键名: goqueue:charactor

    // Push
    countInsert, err := queue.Push(key, "卡布达")

    // Pop
    stringArray, err := queue.Pop(2*time.Second, key)
    
    // GetLength
    intResult, err := queue.GetLength(key)

    // GetRange
    stringArray, err := queue.GetRange(key,0, 9) // 得到数字前10个元素
```