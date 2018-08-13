package system

import (
	"fmt"
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
func LoadConfigInformation(configPath string) (err error) {
	var (
		filePath string
		wr       string
	)

	if configPath == "" {
		wr, _ = os.Getwd()
		wr = path.Join(wr, "conf")

	} else {
		wr = configPath
	}
	common.WorkSpace = wr
	filePath = path.Join(common.WorkSpace, "config.yml")
	fmt.Println(filePath)
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
	// hexiaoyun128 config information
	//common.hexiaoyun128Config = common.ConfigInfo.hexiaoyun128SDK
	// iamge size
	common.Image = common.ConfigInfo.Image
	// quartz task
	common.Quartz = common.ConfigInfo.Quartz

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

	// log config load
	if common.LogInfo.ConfigFile == "" {

		filePath = path.Join(common.WorkSpace, "seelog.xml")
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
	// init api
	if common.InitInfo.Api {
		log.Info("init api data")
		initial_data.InitApi()
	}
	// init system data
	if common.InitInfo.Menu {
		log.Info("init menu data")
		initial_data.InitMenu()
	}
	if common.InitInfo.Role {
		log.Info("init role and policy data")
		initial_data.InitBaseRole()
	}



	var user models.User
	user.Name = common.AdminUserInfo.Name
	user.Email = common.AdminUserInfo.Email
	user.Mobile = common.AdminUserInfo.Mobile
	nn := common.DB
	nn.Where("name = ? AND email = ?", user.Name, user.Email).First(&user)
	if user.ID == 0 {
		user.Password = common.AdminUserInfo.Password
		user.Create(nn)

		user.SetPwd(nn)
		user.IsAdmin = true
		user.AdminAction(nn)
		///*services.UserCreate(&user)
		// 获得管理员角色
		var roleSlice []models.Role
		if admin, e, _ := services.GetRoleByCode("admin"); e == nil {
			roleSlice = append(roleSlice, *admin)
		}
		common.DB.Model(user).Association("Roles").Append(roleSlice)
	}
	middles.InitKeys()

	// hexiaoyun128

	/*if common.hexiaoyun128Config.Enable {
		common.hexiaoyun128Ins = &hexiaoyun128_sdk_golang.hexiaoyun128{}
		common.hexiaoyun128Ins.Config = common.hexiaoyun128Config
		// login
		 common.hexiaoyun128Ins.AutoLogin()


	}*/

	// hexiaoyun128 login get token
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
