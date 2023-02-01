package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080") // Listen函数创建了一个net.Listener的对象，监听一个端口的连接
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		// listener对象的Accept方法会直接阻塞，直到一个新的连接被创建，
		// 然后会返回一个net.Conn对象来表示这个连接
		if err != nil {
			log.Print(err)
			continue
		}
		// handleConn(conn)
		go handleConn(conn)
	}
}

// 处理一个完整的客户端连接
func handleConn(c net.Conn) {
	defer c.Close() // 关闭服务器侧的连接，然后返回到主函数，继续等待下一个连接请求
	for {
		// 直接写入
		_, err := io.WriteString(c, time.Now().Format("08:00:00\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
