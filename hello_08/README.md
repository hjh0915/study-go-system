接口
====
sort.Interface
---------------
内置排序算法接口

http.Handler
------------
例子：将数据库中物品的价格映射成美元
它将库存清单模型化为一个命名为database的map类型，我们给这个类型一个ServeHttp方法，这样它可以满足http.Handler接口。这个handler会遍历整个map并输出物品信息。

由两个不同的web服务器监听不同的端口，并且定义不同的URL将它们指派到不同的handler。我们只要构建另外一个ServeMux并且再调用一次ListenAndServe

error
-----
返回错误信息