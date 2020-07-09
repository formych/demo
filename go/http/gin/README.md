# gin

一个基于`gin`的简单`http`服务

用于验证docker, kubernetes, istio的各种功能的http demo

## 路由

* query
* param

## docker

```sh
# 构建
docker build -t formych/hello:v1
# 运行
docker run -d --name myhello formych/hello:v1
# 查看
docker exec -it myhello /bin/sh
```
