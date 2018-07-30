package system

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/initial_data"
	"github.com/hexiaoyun128/gin-base-framework/middles"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"github.com/cihub/seelog"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

//LoadConfigInformation load config information for application
func LoadConfigInformation() (err error) {
	var (
		filePath string
	)
	wr, _ := os.Getwd()

	filePath = path.Join(wr, "conf", "config.yml")
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		seelog.Errorf(" config file read failed: %s", err)
		seelog.Flush()
		os.Exit(-1)

	}
	err = yaml.Unmarshal(configData, &common.ConfigInfo)
	if err != nil {
		seelog.Errorf(" config parse failed: %s", err)
		seelog.Flush()
		os.Exit(-1)
	}

	// admin user information
	common.AdminUserInfo = common.ConfigInfo.AdminUser
	// server information
	common.ServerInfo = common.ConfigInfo.Server
	// database information
	common.DatabaseInfo = common.ConfigInfo.Database
	// redis information
	common.RedisInfo = common.ConfigInfo.Redis
	// we chat information
	common.WeChatInfo = common.ConfigInfo.WeChat
	// file upload information
	common.FileUploadInfo = common.ConfigInfo.FileUpload
	// qi niu cloud storage information
	common.QiNiuInfo = common.ConfigInfo.QiNiu
	// send message queue information
	common.SendMessageQueueInfo = common.ConfigInfo.SendMessageQueue
	// receive message queue information
	common.ReceiveMessageQueueInfo = common.ConfigInfo.ReceiveMessageQueue
	// we chat information
	common.WeChatInfo = common.ConfigInfo.WeChat
	// file upload information
	common.FileUploadInfo = common.ConfigInfo.FileUpload
	// qi niu cloud storage information
	common.QiNiuInfo = common.ConfigInfo.QiNiu
	// send message queue information
	common.SendMessageQueueInfo = common.ConfigInfo.SendMessageQueue
	// receive message queue information
	common.ReceiveMessageQueueInfo = common.ConfigInfo.ReceiveMessageQueue
	// log information
	common.LogInfo = common.ConfigInfo.Log
	// init information
	common.InitInfo = common.ConfigInfo.Init


	//database connect
	var db *gorm.DB
	switch common.DatabaseInfo.DBType {

	case "postgres":
		db, err = gorm.Open("postgres", common.DatabaseInfo.Connect)
	case "mysql":
		db, err = gorm.Open("mysql", common.DatabaseInfo.Connect)

	}
	if err != nil || db == nil {
		seelog.Errorf("database connect failed: %s", err)
		seelog.Errorf("database config information: %v", common.DatabaseInfo)
		seelog.Flush()
		os.Exit(-1)
	}
	common.DB = db.Set("gorm:save_associations", false)
	// create database tables
	DBMigrate(common.DB)
	//redis connect
	common.RedisClient = redis.NewClient(&redis.Options{
		Addr:     common.RedisInfo.Address,
		Password: common.RedisInfo.Password,
		DB:       common.RedisInfo.DbName,
	})
	//admin user insert to database
	common.AdminUserInfo = common.ConfigInfo.AdminUser
	var user models.User
	user.Name = common.AdminUserInfo.Name
	user.Email = common.AdminUserInfo.Email
	common.DB.Where("name = ? AND email = ?", user.Name, user.Email).First(&user)
	if user.ID == 0 {
		user.Password = common.AdminUserInfo.Password
		user.IsAdmin = true
		services.UserCreate(&user)
	}

	// log config load
	if common.LogInfo.ConfigFile == "" {
		wr, _ := os.Getwd()
		filePath = path.Join(wr, "conf", "seelog.xml")
	} else {
		filePath = path.Join(common.LogInfo.ConfigFile, "seelog.xml")
	}
	log, err := seelog.LoggerFromConfigAsFile(filePath)
	if err != nil {
		seelog.Critical("err parsing see log config file", err)
		seelog.Flush()
		return
	}
	seelog.ReplaceLogger(log)

	// init system data
	if common.InitInfo.Menu {
		log.Info("init menu data")
		initial_data.InitMenu()
	}
	if common.InitInfo.Role {
		log.Info("init role data")
		initial_data.InitBaseRole()
	}

	middles.InitKeys()

	// message queue init
	//message queue init
	//if system.ServerConfig.MqEnable {
	//	amqp.InitAmqpSend()
	//	amqp.InitAmqpReceive()
	//	defer amqp.AmqpReceiveDefer()
	//	defer amqp.AmqpSendDefer()
	//}
	return

}
