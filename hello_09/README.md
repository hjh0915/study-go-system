Goroutines
==========
每一个并发的执行单元叫作一个goroutine,可将goroutine类比作一个线程
clock服务器每一个连接都会起一个goroutine

创建
----
``` go
go f()
```

并发的Clock服务
==============
顺序执行的时钟服务器，它会每隔一秒钟将当前时间写到客户端
需要安装nc工具，也可以用telnet来实现

telnet
======
连接中读取数据，并将读到的内容写到标准输出中，在终端输出。使用了mustCopy函数
``` go 
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```

执行
----
>>> go build hello.go
>>> ./hello &
>>> nc localhost 8080