[TOC]- [关于我](#关于我)
- [关于我](#关于我)
- [echo-scaffolding](#echo-scaffolding)
  - [简介](#简介)
  - [目的及优势](#目的及优势)
  - [技术点对应文档](#技术点对应文档)
  - [职责](#职责)
    - [文件配置](#文件配置)
    - [路由](#路由)
    - [日志](#日志)
    - [请求日志](#请求日志)
# 关于我
一只孤独的饮酒客...

**持续更新中，喜欢的可以点个star。欢迎贡献代码**

github地址: <a href="https://github.com/jeffcail/echo-scaffolding" target="_blank">echo-scaffolding</a>

个人博客地址: <a href="https://blog.caixiaoxin.cn/?p=757" target="_blank">太阳上的雨天</a>

CSDN: <a href="caixiaoxin.blog.csdn.net">太阳上的雨天></a>

知乎: <a href="https://www.zhihu.com/people/cai-cai-55-44-82">太阳上的雨天</a>

# echo-scaffolding



## 简介
Echo best scaffolding.Fllowing me. Let's go
基于echo框架，搭建一个最好用的脚手架。



## 目的及优势

* 快速上手、快速开发、快速交付
* 高性能、高扩展，避免重复造轮子

## 技术点对应文档
* ini: https://github.com/go-ini/ini
* viper: https://github.com/spf13/viper
* Nacos: https://nacos.io/zh-cn/
* Echo: https://echo.labstack.com/
* Log: https://github.com/uber-go/zap
* Cron: https://github.com/robfig/cron
* lumberjack: https://github.com/natefinch/lumberjack
* UUID: https://github.com/google/uuid

## 职责

### 文件配置
* 支持 ini、yaml格式. 分布式配置支持nacos
* 支持多环境配置 dev、pre、prod, 分别为开发环境、预发布环境、生产环境
* 主配置文件为application.yaml
* 如果同时开启了远程配置和本地配置，远程配置优先级高于本地配置
* 远程和本地配置开关必须开启一个
* ini模型配置文件为.example.ini,将.example.ini拷贝为dev.ini、pre.ini、prod、ini即可
* yaml模型配置文件为.example.yaml,将.example.yaml.yaml、pre.yaml、prod、yaml即可

#### USAGE
* nacos
使用nacos远程配置，需要在主文件开启远程配置开关,开启后自动忽略本地配置文件
  
```yaml
Remote: true
```
```yaml
Debug: true
HTTPBind: :8090
Mysql:
  DbDsn: root:123456@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4
  ShowSql: true
Redis:
  RedisAddr: 127.0.0.1:7379
  Password: 123456
  RedisDb: 0
LoggerPath: ./log/echo-scaffolding.log
```
```go
fmt.Println(conf.Config)
```
```shell
go run echo-scaffolding.go -ip 127.0.0.1 -p 7848 -c echo-scaffolding.yml -g echo-scaffolding
```


* ini格式

使用本地ini配置文件,需要关闭远程配置开关.开启读取本地的开关和开启ini配置开关，并且选择使用的环境.默认使用dev环境的配置文件
```yaml
Local: true
ExtFormat: ini
EnvModel: dev
```



```go
fmt.Println(conf.Config)
```

* yaml格式
  
使用本地yaml配置文件,需要关闭远程配置开关.开启读取本地的开关和开启yaml配置开关，并且选择使用的环境，环境默认使用dev环境的配置文件
```yaml
Local: true
ExtFormat: yaml
EnvModel: 
```

```go
fmt.Println(conf.Config)
```

### 路由
单路由模式和路由组模式
* 路由
```go
e.GET("/ping", func(c echo.Context) error {
    return c.JSON(http.StatusOK, "pong...")
})
```

* 路由組
```go
orderGroup := e.Group("/v1/order")
{
    orderGroup.GET("/detail", handlerorder.Detail)
}
```

### 日志
日志支持路径、日志文件大小、保存周期、压缩、时间可配置化。支持达到配置大小，自动压缩。并生成新的日志文件
```go
uber.EchoScaLog.Info("Info logger demo")
uber.EchoScaLog.Info(fmt.Sprintf("Info logger demo :%d", 123))
uber.EchoScaLog.Error("Error logger demo")
var err = errors.New("test error demo")
uber.EchoScaLog.Error(fmt.Sprintf("Error logger demo: %s", "orderno-13546"), zap.Error(err))
```
效果:
```log
2022-11-30 11:07:05	INFO	router/router.go:39	Info logger demo
2022-11-30 11:07:05	INFO	router/router.go:40	Info logger demo :123
2022-11-30 11:07:05	ERROR	router/router.go:41	Error logger demo
2022-11-30 11:07:05	ERROR	router/router.go:43	Error logger demo: orderno-13546	{"error": "test error demo"}
```

### 请求日志
打印请求崩溃堆栈链路信息
采用UUID。为每次请求打上一个唯一ID标识。请求分飞两种，API请求和网页请求。url path意/api/开头为 API请求日志，其他为网页请求日志。 - 前后端合并部署的情况
支持请求日志控制台输出
