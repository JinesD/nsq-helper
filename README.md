# nsq-helper
封装 nsq 的操作方法

### 安装 
```
go get -u -v github.com/JinesD/nsq-helper
```

### 使用
```
package main 

import (
    fmt

    nhelper "github.com/JinesD/nsq-helper"
    nsq "github.com/nsqio/go-nsq"
)

func main () {
    cfg := nsq.NewConfig()

    nh := &nhelper.NSQHandler{
        Addr: "127.0.0.1:27017",
        Cfg:  cfg,
    }

    if err := nh.Publish("topic", "value"); err != nil {
        fmt.Println("failed to publish topic")
    } 
}
```