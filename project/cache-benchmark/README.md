
cache-benchmark 主要对 cache 项目进行压力测试
==============

## 依赖

> go get github.com/go-redis/redis

## 运行方式

./cache-benchmark -type http -n 100000 -r 100000 -t set

./cache-benchmark -type http -n 100000 -r 100000 -t get