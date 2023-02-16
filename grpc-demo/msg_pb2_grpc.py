# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import msg_pb2 as msg__pb2


class MsgServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetMsg = channel.unary_unary(
                '/MsgService/GetMsg',
                request_serializer=msg__pb2.MsgRequest.SerializeToString,
                response_deserializer=msg__pb2.MsgResponse.FromString,
                )
        self.SimpleFun = channel.unary_unary(
                '/MsgService/SimpleFun',
                request_serializer=msg__pb2.MsgRequest.SerializeToString,
                response_deserializer=msg__pb2.MsgResponse.FromString,
                )


class MsgServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetMsg(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SimpleFun(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_MsgServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetMsg': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMsg,
                    request_deserializer=msg__pb2.MsgRequest.FromString,
                    response_serializer=msg__pb2.MsgResponse.SerializeToString,
            ),
            'SimpleFun': grpc.unary_unary_rpc_method_handler(
                    servicer.SimpleFun,
                    request_deserializer=msg__pb2.MsgRequest.FromString,
                    response_serializer=msg__pb2.MsgResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'MsgService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class MsgService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetMsg(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/MsgService/GetMsg',
            msg__pb2.MsgRequest.SerializeToString,
            msg__pb2.MsgResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def SimpleFun(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/MsgService/SimpleFun',
            msg__pb2.MsgRequest.SerializeToString,
            msg__pb2.MsgResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
