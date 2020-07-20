# Python GRPC

### base install
- grpcio依赖 
    - `pip3 install grpcio`  build非常慢
- py protoc编译器
    - `pip3 install grpcio-tools`
- 编译protoc文件
    - `python -m grpc_tools.protoc --python_out=. --grpc_python_out=. -I. hello.proto`
``` 
python -m grpc_tools.protoc: python 下的 protoc 编译器通过 python 模块(module) 实现, 所以说这一步非常省心
--python_out=. : 编译生成处理 protobuf 相关的代码的路径, 这里生成到当前目录
--grpc_python_out=. : 编译生成处理 grpc 相关的代码的路径, 这里生成到当前目录
-I. helloworld.proto : proto 文件的路径, 这里的 proto 文件在当前目录
```

### 目录结构
- hello.proto
- hello_pb2.py           # protobuf 数据交互
- hello_pb2_grpc.py      # grpc 交互
- hello_grpc_server.py   # grpc server
- hello_grpc_client.py   # grpc client