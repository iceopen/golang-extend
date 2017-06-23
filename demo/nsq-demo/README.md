# 监控HTTP数据

## 依赖
- 必须：https://github.com/shirou/gopsutil
- window必须：https://github.com/go-ole/go-ole
- window必须：https://github.com/StackExchange/wmi

## 编译方式
- go 1.5 一下使用：gox 多平台编译
- go 1.5 以后使用：GOOS=linux GOARCH=amd64 go build -o ./bin/nsqProducer nsq_producer.go 进行编译

## 编译平台参考

可用的OS和ARCH的值如下：

```
$GOOS	$GOARCH
darwin	386
darwin	amd64
darwin	arm
darwin	arm64
dragonfly	amd64
freebsd	386
freebsd	amd64
freebsd	arm
linux	386
linux	amd64
linux	arm
linux	arm64
linux	ppc64
linux	ppc64le
netbsd	386
netbsd	amd64
netbsd	arm
openbsd	386
openbsd	amd64
openbsd	arm
plan9	386
plan9	amd64
solaris	amd64
windows	386
windows	amd64
```
## 部署方案：
- 1.防火墙开发TCP：8899 端口
- 2.Supervisord 进行部署