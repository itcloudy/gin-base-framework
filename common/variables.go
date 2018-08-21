package common

import (
	"github.com/casbin/casbin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	WorkSpace               string             //config
	DB                      *gorm.DB           // database contains
	RedisClient             *redis.Client      // redis client
	Enforcer                *casbin.Enforcer   // casbin
	CasbinRoleIds           []int              // casbin role ids 由于初始化了角色，需要在路由启动之后再将策略加到casbin_rule中
	ConfigInfo              *configModel       // all server config information
	ServerInfo              *serverModel       // server config information
	DatabaseInfo            *databaseModel     // database config information
	RedisInfo               *redisModel        // redis config information
	AdminUserInfo           *adminUserModel    // admin user config information
	WeChatInfo              *weChatModel       // we chat config information
	FileUploadInfo          *fileUploadModel   // file upload config information
	QiNiuInfo               *qiNiuModel        // qi niu cloud storage config information
	SendMessageQueueInfo    *messageQueueModel // send message queue config information
	ReceiveMessageQueueInfo *messageQueueModel // receive message queue config information
	LogInfo                 *logModel          // log config information
	InitInfo                *initModel         // init data config information
	Image                   *image             // image size
	Quartz                  *quartz            // cron task
	Logger                  *zap.Logger        // log instance

)
