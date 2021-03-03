## 架构

- 网关层
- 聚合业务层
- 基础服务层

## 特点

### 服务注册与发现

### 链路追踪

### 熔断（hystrix-go）

#### 作用

解决服务雪崩。

- 阻止故障连锁反应
- 快速失败并迅速恢复
- 然后回退并优雅降级
- 提供近实时的监控和告警

#### 原则

- 防止单独的依赖消耗资源
- 过载快速切断并快速失败，防止排队
- 尽可能回退以保护用户免受故障
- 通过接近实时的指标，监控和告警，确保故障即时发现

#### 原理

| 状态              | 解释                                                         |
| ----------------- | ------------------------------------------------------------ |
| CLOSED关闭状态    | 允许流量通过                                                 |
| OPEN打开状态      | 不允许流量通过，处于降级状态，走降级流程                     |
| HALF_OPEN半开状态 | 允许某些流量通过，如果出现超时、异常，将进入OPEN状态，如果成功，将进入CLOSED状态 |



#### 安装

### 限流

### 负载均衡

## 部署

micro/config/mysql
```
{
    "host":"127.0.0.1",
    "port":3306,
    "username":"root",
    "password":"123456",
    "db":"12306"
}
```
micro/config/redis
```
host = 127.0.0.1
port = 6379
db = 0
```
