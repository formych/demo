# gRPC

通用的RPC框架，将mobile和HTTP/2放在首位

## protoc的插件

[文章](https://developers.google.com/protocol-buffers/docs/reference/cpp/google.protobuf.compiler.plugin)

protoc --NAME_out=plugins=grpc, import_path=mypackage:. *.proto

这里通过 --NAME_out ，就能知道是需要找 protoc-gen-NAME 插件， 即 protoc-gen-go 插件。

--go_out=plugins=grpc 在https://github.com/golang/protobuf

## FAQ

1. 同时支持REST/JSON和gRPC的实现？

* 同时开启2个端口
* [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
