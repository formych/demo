# go-proto

学习，验证protobuf的go版本

## 编译proto

```sh
prtoc -I=$SRC_DIR --go_out=$DEST_DIR $SRC_DIR/*.proto
```

## 新旧版本

新旧版本不兼容

### v1

github.com/golang/protobuf为旧版本的protobuf实现，官方v1为1.3.5， [1.4.0, 1.20.0)为基于v2实现
[github](https://github.com/golang/protobuf)

### v2

google.golang.org/protobuf为新版本实现，[1.20.0, ...)

[github](https://github.com/protocolbuffers/protobuf-go)
