from concurrent import futures
import time
import grpc
import hello_pb2
import hello_pb2_grpc

# 实现proto Greeter
class Greeter(hello_pb2_grpc.GreeterServicer):
    # proto grpc
    def SayHello(self,request,context):
        time.sleep(60 * 10)
        return hello_pb2.Response(message = 'hello {}'.format(request.name))

def server():
    # run grpc
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))  # 开启异步 设定线程池大小
    hello_pb2_grpc.add_GreeterServicer_to_server(Greeter(),server)
    server.add_insecure_port('[::]:8089')
    server.start()
    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == "__main__":
    server()