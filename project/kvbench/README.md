# KVBench

压力测试案例

## 构建方法

```
make
```

## 案例

启动服务

```
 ./kvbench --store=badger --path=badger.db 
 ```
 
 使用 redis-benchmark 进行压力测试 
 
 
 ```
 redis-benchmark -p 6380 -q -t set,get
 ```
 
 得到的结果
 
 ```
 SET: 39308.18 requests per second
 GET: 67204.30 requests per second
 ```