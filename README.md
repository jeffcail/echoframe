[TOC]- [关于我](#关于我)
- [关于我](#关于我)
- [echo-scaffolding](#echo-scaffolding)
  - [简介](#简介)
  - [安装](#安装)
  - [目的及优势](#目的及优势)
  - [技术点对应文档](#技术点对应文档)
  - [职责](#职责)
    - [1. 文件配置](#文件配置)
    - [2. 路由](#路由)
    - [3. 日志](#日志)
    - [4. 请求日志](#请求日志)
    - [5. GOMAXPROCS](#GOMAXPROCS)
    - [6. HTTP请求](#HTTP请求)
    - [7. API统一格式返回支持jsonxml](#API统一格式返回支持jsonxml)
    - [8. 参数校验](#参数校验)  
    - [9. 数据库支持mysql](#数据库支持mysql)  
    - [10. XORM](#XORM)
    - [11. 根据数据库自动生成gostruct](#根据数据库自动生成gostruct)
    - [12. Redis](#Redis)
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

## 安装
安装完之后名字echo-scaffolding可改，可根据自己的需求精简或者添加架子结构。也可直接使用
```shell
git clone https://github.com/jeffcail/echo-scaffolding.git

cd echo-scaffolding

go mod tidy

go run echo-scaffolding
```

## 目的及优势

* 快速上手、快速开发、快速交付
* 高性能、高扩展，避免重复造轮子

## 技术栈和对应的包
* ini: https://github.com/go-ini/ini
* viper: https://github.com/spf13/viper
* Nacos: https://nacos.io/zh-cn/
* Echo: https://echo.labstack.com/
* Log: https://github.com/uber-go/zap
* Cron: https://github.com/robfig/cron (github.com/robfig/cron v1.2.0)
* lumberjack: https://github.com/natefinch/lumberjack
* UUID: https://github.com/google/uuid
* gorequest: https://github.com/jeffcail/gorequest
* mysql: github.com/go-sql-driver/mysql  
* xorm: https://github.com/go-xorm/xorm
* Redis: gopkg.in/redis.v5

## 职责

### 文件配置
* 支持 ini、yaml格式. 分布式配置支持nacos
* 支持多环境配置 dev、pre、prod, 分别为开发环境、预发布环境、生产环境
* 主配置文件为application.yaml
* 如果同时开启了远程配置和本地配置，远程配置优先级高于本地配置
* 远程和本地配置开关必须开启一个
* ini模型配置文件为.example.ini,将.example.ini拷贝为dev.ini、pre.ini、prod、ini即可
* yaml模型配置文件为.example.yaml,将.example.yaml.yaml、pre.yaml、prod、yaml即可

#### 用法
* nacos
使用nacos远程配置，需要在主文件开启远程配置开关,开启后自动忽略本地配置文件
  
```yaml
Remote: true
```
nacos配置案例:
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

使用本地ini配置文件,需要关闭远程配置开关.开启读取本地的开关和开启ini配置开关，并且选择使用的环境.不选择默认使用dev环境的配置文件
```yaml
Local: true
ExtFormat: ini
EnvModel: 
```



```go
fmt.Println(conf.Config)
```

* yaml格式
  
使用本地yaml配置文件,需要关闭远程配置开关.开启读取本地的开关和开启yaml配置开关，并且选择使用的环境，不选择环境默认使用dev环境的配置文件
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

### GOMAXPROCS
主配置文件增加是否开启GOMAXPROCS开关. 默认不开启
```json
IsEnableGOMAXPROCS: false
```

### HTTP请求
集成了HTTP请求.支持常用的请求方式GET、POST.具体使用详情查看<a href="https://github.com/jeffcail/gorequest" target="_blank">gorequest</a>

### API统一格式返回支持jsonxml
json
```go
return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
```

xml
```go
return utils.ToXml(c, utils.Res.Response(false, msg, code.FAILED))
```

### 参数校验
用法
```go
// CreateUserInput
type CreateUserInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=12"`
}

// CreateUser
func CreateUser(c echo.Context) error {
    param := &CreateUserInput{}
    _ = c.Bind(param)
    msg := utils.ValidateParam(param)
    if msg != "" {
    return utils.ToJson(c, utils.Res.Response(false, msg, code.FAILED))
    }
    return nil
}

```


### 数据库支持mysql
使用
```go
db.Mysql.Table()
```

开启事物
```go
err := db.Mysql.Transaction(func(s *xorm.Session) error {
    insert, err := s.Insert(data)
    if err != nil {
    return err
    }
    if insert != 1 {
    return errors.New("data write failed")
    }
    
    update, err := s.Update(data)
    if err != nil {
    return err
    }
    if update != 1 {
    return errors.New("data update failed")
    }
    return nil
})
if err != nil {
log.Println(err)
```

### 根据数据库自动生成gostruct
```shell
go get github.com/go-xorm/cmd/xorm
go get github.com/go-xorm/xorm
```
到GOPATH\src\github.com\go-xorm\cmd\xorm 目录下，执行

```shell
go build
```

用法
修改脚本里面的数据连接信息和数据库名字
```shell
./sqlToStruct.sh
```

### Redis
用法
```go
db.Rc.Set("a", 10, -1)
```
