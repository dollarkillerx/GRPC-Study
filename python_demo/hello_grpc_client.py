import grpc
import hello_pb2
import hello_pb2_grpc

def run():
    # connect grpc server
    channel = grpc.insecure_channel('localhost:8089')
    # send grpc
    stub = hello_pb2_grpc.GreeterStub(channel)

    response = stub.SayHello(hello_pb2.Request(name = 'cpx'))
    print("Greeter client received: " + response.message)

if __name__ == '__main__':
    run()