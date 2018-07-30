package common

import (
	"github.com/casbin/casbin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var (
	DB                      *gorm.DB                          // database contains
	RedisClient             *redis.Client                     // redis client
	Enforcer                *casbin.Enforcer                  // casbin
	ConfigInfo              *configModel                      // all server config information
	ServerInfo              *serverModel                      // server config information
	DatabaseInfo            *databaseModel                    // database config information
	RedisInfo               *redisModel                       // redis config information
	AdminUserInfo           *adminUserModel                   // admin user config information
	WeChatInfo              *weChatModel                      // we chat config information
	FileUploadInfo          *fileUploadModel                  // file upload config information
	QiNiuInfo               *qiNiuModel                       // qi niu cloud storage config information
	SendMessageQueueInfo    *messageQueueModel                // send message queue config information
	ReceiveMessageQueueInfo *messageQueueModel                // receive message queue config information
	LogInfo                 *logModel                         // log config information
	InitInfo                *initModel                        // init data config information

)
