并发获取多个URL
=============
main函数中用make函数创建了一个传递string类型参数的channel，对每一个命令行参数，我们都用go这个关键字来创建一个goroutine，并且让函数在这个goroutine异步执行http.Get方法。
用main函数来接收所有fetch函数传回的字符串，可以避免在goroutine异步执行还没有完成时main函数提前退出。

goroutine
---------
函数的并发执行方式
go function则表示创建一个新的goroutine，并在这个新的goroutine中执行这个函数。

channel
-------
在goroutine之间进行参数传递

执行
====
>>> go build hello.go
>>> ./hello url1 url2...