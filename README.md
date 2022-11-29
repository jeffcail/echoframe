[TOC]- [关于我](#关于我)
- [关于我](#关于我)
- [echo-scaffolding](#echo-scaffolding)
  - [简介](#简介)
  - [目的及优势](#目的及优势)
  - [技术点对应文档](#技术点对应文档)
  - [职责](#职责)
    - [文件配置](#文件配置)
    - [路由](#路由)
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

## 职责

### 文件配置
支持 ini、yaml格式. 分布式配置支持nacos
#### USAGE
* ini格式
```go
func main() {
	boot.Bootstrap()
    fmt.Println(confini.Config())
}
```

* yaml格式
```go
func main() {
	boot.Bootstrap()
    fmt.Println(confyaml.YConf)
}
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
func main() {
	boot.Bootstrap()
    fmt.Println(confnacos.NConfig())
}
```
```shell
go run echo-scaffolding.go -ip 127.0.0.1 -p 7848 -c echo-scaffolding.yml -g echo-scaffolding
```