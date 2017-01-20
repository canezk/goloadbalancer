# 负载均衡
```
关键词：golang,zookeeper,一致性哈希,nio
```
## 设计初衷
### 出发点
```
1、golang高并发，1.7之后的gc性能提升，很适合作为一个中间件转发
2、微服务架构盛行，zookeeper（或者类似的接口）使用很多，希望实现一款通用的
   借助分布式coodinator做负载均衡的中间件
3、适合非强一致性的服务，在设计的时候，设置超时，如果从zk上面读取到的节点响应
    很慢，那么立刻切换到下一个节点

```
### 缺点
```
1、引入一款中间件无疑就是提高了系统复杂度，提高了故障概率
2、golang相关的常用组件可能会不够成熟相对于java而言
```

### 主要目标
```
1、动态负载均衡
2、对现有数据流透明
```

### 数据流
```
1、给下游服务提供api，注册（域名，znode信息）
2、收到url之后，解析url，获取域名，读取对应的znode server信息
3、一致性哈希url，然后请求对应的server
```

## 主要模块
### client的connector
### zookeeper的connector
```
负责和zookeeper通信，注册watcher
```
#### zookeeper golang client选择
```
gozk：https://wiki.ubuntu.com/gozk
go-zookeeper：https://github.com/samuel/go-zookeeper
后者文档和代码都是开源的
```
### 一致性哈希
```
根据从zookeeper上面读取的节点信息，进行一致性哈希
https://github.com/serialx/hashring 
```
### 降级策略
```
1、和zookeeper的通信出现问题？缓存上一次读取的节点
2、应用层的策略需要应用服务自己考虑？（暂时没想到好的办法）
```
