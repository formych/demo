# pprof

## 使用

+ profile
+ head

### 方案1

```sh
go tool pprof -http=:32000 http://127.0.0.1:6060/debug/pprof/heap
```

浏览器直接访问就可以

### 方案2

```sh
go tool pprof ./binary http://127.0.0.1:6060/debug/pprof/heap
```

进入，top查看最高的，执行web保存对应svg图片，使用浏览器打开图片，进入ui界面
