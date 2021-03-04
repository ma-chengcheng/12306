## 架构

- 网关层
- 聚合业务层
- 基础服务层

## 特点

### 服务注册与发现

### 链路追踪

### 熔断（hystrix-go）

### 限流

### 负载均衡

## 

## 部署

micro/config/mysql
```
{
    "host":"0.0.0.0",
    "port":3306,
    "username":"root",
    "password":"123456",
    "db":"12306"
}
```
micro/config/redis
```
{
    "host":"0.0.0.0",
    "port":6379,
    "db":0
}
```
