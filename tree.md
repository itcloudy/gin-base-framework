
```
├── Dockerfile                      // docker file
├── README.md                       // 说明文件
├── common                          // 通用文件
│   ├── common.go                   // 通用函数和响应封装函数
│   ├── config.go                   // 系统配置对应的struct，yaml获得配置 
│   ├── const.go                    // 系统常量
│   ├── response_message.go         // 响应信息转换
│   └── variables.go                // 系统全局变量
├── conf                            // 配置文件夹
│   ├── api_data.yml                // 系统所有的对外接口，将会在系统第一次启动的时候初始化到数据库
│   ├── casbin_rbac_model.conf      // casbin 策略文件，用于权限控制
│   ├── config.yml                  // 系统配置yaml
│   ├── https                       // https 公私钥
│   │   ├── cert.pem
│   │   └── key.pem
│   ├── jwt                         // jwt 公私钥
│   │   ├── app.rsa
│   │   ├── app.rsa.pub
│   │   ├── tm.rsa
│   │   └── tm.rsa.pub  
│   ├── menu_data.yml               // 系统菜单，前后端分类从后台返回用户角色对应的菜单，
│   └── role_data.yml               // 角色对应的菜单和接口初始化数据，将会在系统第一次启动的时候初始化到数据库
├── controllers                     // API请求对应的controller
│   ├── base_index.go               // 测试和token刷新
│   ├── base_login.go               // 手机号和微信登录
│   ├── base_menu.go                // 系统菜单相关controller
│   ├── base_mq.go                  // 消息队列测试请求
│   ├── base_role.go                // 角色相关的controller
│   ├── base_upload.go              // 文件上传
│   ├── base_user.go                // 用户相关controller
│   ├── base_verify_code.go         // 验证码，暂未支持
│   └── third_wechat.go             // 微信相关
├── crons                           // 后台定时任务
│   └── start.go
├── daemons                         // casbin权限对数据操作
│   └── enforcer.go
├── docs                            // swagger 文档
│   ├── docs.go
│   └── swagger
│       ├── swagger.json
│       └── swagger.yaml
├── elastic                         // 搜索引擎，暂未支持
├── initial_data                    // 数据初始化函数
│   ├── api_data.go                 // 系统接口
│   ├── base_role.go                // 系统角色
│   └── menu_data.go                // 系统菜单
├── logs                            // 日志
│   └── gin-base-framework.log
├── main.go                         // 应用入口
├── middles                         // 中间件
│   ├── auth.go                     // 权限验证中间件
│   ├── cors.go                     // 跨域中间件
│   ├── jwt.go                      // token中间件
│   └── visit.go                    // 访问记录中间件`
├── models                          // 数据库对应的models
│   ├── base.go                     // 基础model，其他model继承该model
│   ├── base_api.go                 // 接口
│   ├── base_menu.go                // 菜单
│   ├── base_role.go                // 角色
│   ├── base_role_api.go            // 角色拥有的API
│   ├── base_role_menu.go           // 角色拥有的菜单
│   ├── base_user.go                // 用户
│   ├── wx_account.go               // 微信openid
│   ├── wx_app_pay_info.go          // 微信支付记录
│   └── wx_pay_resp.go              // 微信支付响应
├── mqs                             // 消息队列
│   ├── amqp                        // amqp 已实现
│   │   ├── readme.md
│   │   ├── receive.go
│   │   └── send.go
│   └── kafka                       // kafka 暂未实现
│       └── readme.md
├── router                          // 路由
│   ├── add_router.go               // api 
│   └── router.go                   // 中间件
├── services
│   ├── base_api.go
│   ├── base_menu.go
│   ├── base_role.go
│   ├── base_role_api.go
│   ├── base_role_menu.go
│   ├── base_user.go
│   └── wx_account.go
├── storage                         // 文件存储
│   ├── local.go                    // 本地
│   └── qiniu.go                    // 七牛
├── system                          // 系统配置
│   ├── config.go                   // 系统配置
│   └── db_migrate.go               // 数据库 merge
├── system_statics                  // 系统自带文件，部署时添加
├── tree.md
├── upload                          // 用户文件上传路径，本地存储时，文件名为内容hash
│   └── 7ff70f285fd232d8a0e2942dac63cc84.jpeg
├── utils                           // 第三方函数    
│   └── validate.go                 // 验证函数
└── vendor                          // 依赖
    └── vendor.json

140 directories, 667 files
