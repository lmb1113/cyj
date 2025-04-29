## 客户端编译启动
```l
wails build -platform=windows,linux,darwin -upx  -ldflags "-s -w"
```

## 服务端启动
```azure
go run server/main.go
```

## 服务端编译
```azure
go build server/main.go
```

![img.png](docs/images/home.png)
