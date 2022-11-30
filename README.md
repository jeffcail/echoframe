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

## 职责

### 文件配置
支持 ini、yaml格式. 分布式配置支持nacos
#### USAGE
* ini格式
```go
fmt.Println(confini.Config())
```

* yaml格式
```go
fmt.Println(confyaml.YConf)
```

* nacos
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
fmt.Println(confnacos.NConfig())
```
```shell
go run echo-scaffolding.go -ip 127.0.0.1 -p 7848 -c echo-scaffolding.yml -g echo-scaffolding
```

### 路由
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
