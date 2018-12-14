# Gin Base Framework

## 建议使用新框架 [https://github.com/itcloudy/base-framework/tree/2.0](https://github.com/itcloudy/base-framework/tree/2.0)
## [目录说明](tree.md) 

将`conf/config_bak.yml`修改为`conf/config.yml`
请在配置文件: `conf/config.yml`中修改数据链接和端口说明
运行程序:
```go
go build
./gin-base-framework
```
## 完成的功能
* jwt认证
* casbin权限控制,支持角色集成，根据角色获得菜单
* 微信openid获得
* 支持mysql，postgresql，sqlite数据库
* rabbit消息队列
* 微信小程序支付

## 备注
已注释掉 `authRouter.Use(middles.CasbinJwtAuthorize(common.Enforcer))`
权限开启请恢复

## Swagger Docs
* 增加路由`	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))`
* swagger文档生成
```sh
swag init
```
* 启动应用
* 打开浏览器[http://127.0.0.1:8000/swagger/index.html](http://127.0.0.1:8000/swagger/index.html)

## Docker 镜像
```sh 
docker build -t gin-base-framework .
```
## 功能说明及相关文档
- [Elasticsearch部署](elasticsearch.md)
## QQ群
![](QQ20180913-0.jpg)
