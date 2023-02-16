import grpc
 
import msg_pb2
import msg_pb2_grpc
 
def run():
    # 连接 rpc 服务器
    with grpc.insecure_channel('localhost:50051') as channel:
        # 调用 rpc 服务
        client = msg_pb2_grpc.MsgServiceStub(channel)
        response1 = client.GetMsg(msg_pb2.MsgRequest(name='world'))
        response2 = client.SimpleFun(msg_pb2.MsgRequest(name='name'))
    print("Client received: " + response1.msg)
    print("Client received: " + response2.msg)

 
if __name__ == '__main__':
    run()