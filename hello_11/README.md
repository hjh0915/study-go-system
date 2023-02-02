channel创建
===========
``` go
ch := make(chan int) 
```

channel操作
===========
发送和接收两个操作都使用<-运算符

发送
----
>>> ch <- x

接收
----
>>> x = <-ch

关闭
----
>>> close(ch)

无缓存channel
============
调用make函数创建的是一个无缓存的channel，同步Channels
>>> ch = make(chan int)
>>> ch = make(chan int, 0)

单方向Channel
============
拆分函数
``` go
func counter(out chan int)
func squarer(out, in chan int)
func printer(in chan int)
```

模块分离
=======
Go Module

> $ go mod init hello_11

生成 go.mod  

    module hello_11

    go 1.19.5

golang第三方包引用报错
====================
>>> go env -w GO111MODULE=auto