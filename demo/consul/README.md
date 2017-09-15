Consul 案例使用说明
==============

## 步奏
* 1.启动 consul  参考命令 consul agent -dev -ui -node=consul-dev -client=192.168.1.8 （参考具体参照官方）
* 2.go run server/main.go
* 3.go run client/main.go

## 依赖
> go get github.com/hashicorp/consul/api

## 细节说明
通过
