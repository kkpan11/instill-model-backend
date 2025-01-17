# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import ray_pb2 as ray__pb2


class RayServiceStub(object):
    """Ray service for internal process
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ModelReady = channel.unary_unary(
                '/ray.serve.RayService/ModelReady',
                request_serializer=ray__pb2.ModelReadyRequest.SerializeToString,
                response_deserializer=ray__pb2.ModelReadyResponse.FromString,
                )
        self.ModelMetadata = channel.unary_unary(
                '/ray.serve.RayService/ModelMetadata',
                request_serializer=ray__pb2.ModelMetadataRequest.SerializeToString,
                response_deserializer=ray__pb2.ModelMetadataResponse.FromString,
                )
        self.ModelInfer = channel.unary_unary(
                '/ray.serve.RayService/ModelInfer',
                request_serializer=ray__pb2.ModelInferRequest.SerializeToString,
                response_deserializer=ray__pb2.ModelInferResponse.FromString,
                )


class RayServiceServicer(object):
    """Ray service for internal process
    """

    def ModelReady(self, request, context):
        """ModelReady method receives a ModelReadyRequest message and
        returns a ModelReadyResponse
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ModelMetadata(self, request, context):
        """ModelMetadata method receives a ModelMetadataRequest message and
        returns a ModelMetadataResponse
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ModelInfer(self, request, context):
        """ModelInfer method receives a ModelInferRequest message and
        returns a ModelInferResponse
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_RayServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ModelReady': grpc.unary_unary_rpc_method_handler(
                    servicer.ModelReady,
                    request_deserializer=ray__pb2.ModelReadyRequest.FromString,
                    response_serializer=ray__pb2.ModelReadyResponse.SerializeToString,
            ),
            'ModelMetadata': grpc.unary_unary_rpc_method_handler(
                    servicer.ModelMetadata,
                    request_deserializer=ray__pb2.ModelMetadataRequest.FromString,
                    response_serializer=ray__pb2.ModelMetadataResponse.SerializeToString,
            ),
            'ModelInfer': grpc.unary_unary_rpc_method_handler(
                    servicer.ModelInfer,
                    request_deserializer=ray__pb2.ModelInferRequest.FromString,
                    response_serializer=ray__pb2.ModelInferResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'ray.serve.RayService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class RayService(object):
    """Ray service for internal process
    """

    @staticmethod
    def ModelReady(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ray.serve.RayService/ModelReady',
            ray__pb2.ModelReadyRequest.SerializeToString,
            ray__pb2.ModelReadyResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ModelMetadata(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ray.serve.RayService/ModelMetadata',
            ray__pb2.ModelMetadataRequest.SerializeToString,
            ray__pb2.ModelMetadataResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ModelInfer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/ray.serve.RayService/ModelInfer',
            ray__pb2.ModelInferRequest.SerializeToString,
            ray__pb2.ModelInferResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
