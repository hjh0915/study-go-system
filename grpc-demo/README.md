RPC
===
RPC远程过程调用，一种通过网络从远程计算机程序上请求服务。采用客户端/服务器模式。客户端请求程序，服务器提供服务。

gRPC
====
客户端应用程序可以直接调用不同机器上的服务器应用程序上的方法，就像本地对象一样，更容易创建分布式应用程序和服务。gRPC基于定义服务的想法，指定了可以用参数远程调用的方法和返回类型。

使用原理
=======
服务定义从.proto开始，gRPC 提供生成客户端和服务器端代码的插件。gRPC 用户通常在客户端调用这些 API，并在服务器端实现相应的 API。

在服务器方面，服务器实现服务声明的方法，并运行 gRPC 服务器来处理客户端请求。gRPC 基础框架解码传入的请求、执行服务方法并编码服务响应。

在客户端方面，客户端有一个称为stub的本地对象（对于某些语言，首选术语是客户端），该对象实现与服务相同的方法。然后，客户端只需在本地对象上调用这些方法，在适当的protbuf消息类型中包装请求的参数 - gRPC 在将请求发送到服务器并返回服务器的协议缓冲响应后进行处理。

proto
=====
在编写好helloworld.proto文件后，切换至proto目录执行：
> $ python3 -m grpc_tools.protoc -I./ --python_out=. --grpc_python_out=. ./helloworld.proto
*自动生成helloworld_pb2_grpc.py和helloworld_pb2.py两个文件*

创建服务端
=========
在 Server 中，主要是实现服务，按照 msg.proto 定义的，这里需要写一个服务类 MsgServicer ，这个类需要实现之前定义的 GetMsg。GetMsg 接收到的请求是在 request 中， msg.proto 中定义的 name 就是 request.name ，接着在 GetMsg 中设计 msg.proto中定义的 MsgResponse。

创建客户端
=========
使用 grpc.insecure_channel('localhost:50051') 进行连接 服务端， 接着在这个 channel 上创建 stub ， 在 msg_pb2_grpc 里可以找到 MsgServiceStub 这个类相关信息。这个 stub 可以调用远程的 GetMsg 函数。 MsgRequest 中的 name 即 msg.proto 中定义的数据。在回应里可以得到 msg.proto 中定义的 msg 。

运行
====
运行服务器，它会监听50051端口：
> $ python msg_server.py

在另一个终端运行客户端：
> $ python msg_guide_client.py