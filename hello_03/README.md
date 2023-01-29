获取一个URL
==========
http.Get函数是创建HTTP请求的函数
ioutil.ReadAll函数从response中读取到全部内容；将其结果保存在变量b中。

nil
====
空值

执行
====
>>> go build hello.go
>>> ./hello http://gopl.io