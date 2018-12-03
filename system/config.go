package system

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/go-redis/redis"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/itcloudy/gin-base-framework/initial_data"
	"github.com/itcloudy/gin-base-framework/middles"
	"github.com/itcloudy/gin-base-framework/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/yaml.v2"
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
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)

	}
	err = yaml.Unmarshal(configData, &common.ConfigInfo)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)

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
	// itcloudy config information
	//common.itcloudyConfig = common.ConfigInfo.itcloudySDK
	// iamge size
	common.Image = common.ConfigInfo.Image
	// quartz task
	common.Quartz = common.ConfigInfo.Quartz

	// log config load
	if common.LogInfo.Mode == "basic" {
		logPath := common.LogInfo.Path   //log path
		logLevel := common.LogInfo.Level // log level
		isDebug := true                  // log mode
		if common.ServerInfo.Mode == "release" {
			isDebug = false
		}
		initBasicLogger(logLevel, logPath, isDebug)
		log.SetFlags(log.Lmicroseconds | log.Lshortfile | log.LstdFlags)
	} else if common.LogInfo.Mode == "advanced" {
		initAdvancedLogger()
	}

	// elastic
	if common.ConfigInfo.Elastic.Enable {
		// elasticsearch need auth
		if common.ConfigInfo.Elastic.Auth {
			common.ElasticClient, err = elastic.NewClient(
				elastic.SetSniff(common.ConfigInfo.Elastic.SnifferEnabled),
				elastic.SetURL(common.ConfigInfo.Elastic.ServerAddress...),
				elastic.SetBasicAuth(common.ConfigInfo.Elastic.AuthUsername, common.ConfigInfo.Elastic.AuthPassword),
			)
		} else {
			common.ElasticClient, err = elastic.NewClient(
				elastic.SetSniff(common.ConfigInfo.Elastic.SnifferEnabled),
				elastic.SetURL(common.ConfigInfo.Elastic.ServerAddress...),
			)
		}
		if err != nil {
			common.Logger.Fatal("elasticsearch connect failed " + err.Error())
		}
		if info, code, err := common.ElasticClient.Ping(common.ConfigInfo.Elastic.ServerAddress[0]).Do(context.Background()); err != nil {
			common.Logger.Fatal("elasticsearch ping failed " + err.Error())
		} else {
			common.Logger.Info(fmt.Sprintf("elasticsearch name: %s version: %+v code: %d", info.Name, info.Version, code))
		}

	}

	//database connect
	var db *gorm.DB
	switch common.DatabaseInfo.DBType {

	case "postgres":
		db, err = gorm.Open("postgres", common.DatabaseInfo.Connect)
	case "mysql":
		db, err = gorm.Open("mysql", common.DatabaseInfo.Connect)

	}
	if err != nil || db == nil {
		fmt.Printf("database connect failed: %s", err)
		fmt.Printf("database config information: %v", common.DatabaseInfo)
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

	// init api
	if common.InitInfo.Api {
		initial_data.InitApi()
	}
	// init system data
	if common.InitInfo.Menu {
		initial_data.InitMenu()
	}
	if common.InitInfo.Role {
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
	}
	middles.InitKeys()

	// message queue init
	//if system.ServerConfig.MqEnable {
	//	amqp.InitAmqpSend()
	//	amqp.InitAmqpReceive()
	//	defer amqp.AmqpReceiveDefer()
	//	defer amqp.AmqpSendDefer()
	//}
	return

}
func initBasicLogger(logLevel string, logPath string, isDebug bool) {
	if _, err := os.Open(logPath); err != nil && os.IsNotExist(err) {
		p, _ := os.Getwd()
		logPath = path.Join(p, "logs")
		os.Mkdir(logPath, os.ModePerm)
		logPath = path.Join(logPath, "gin-base-framework.log")
		os.Create(logPath)
	} else {
		os.Remove(logPath)
		os.Create(logPath)
	}

	var js string
	if isDebug {
		js = fmt.Sprintf(`{
              "level": "%s",
              "encoding": "json",
              "outputPaths": ["stdout","%s"],
              "errorOutputPaths": ["stdout"]
              }`, logLevel, logPath)
	} else {
		js = fmt.Sprintf(`{
              "level": "%s",
              "encoding": "json",
              "outputPaths": ["%s"],
              "errorOutputPaths": ["%s"]
              }`, logLevel, logPath, logPath)
	}

	var cfg zap.Config
	if err := json.Unmarshal([]byte(js), &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var err error
	common.Logger, err = cfg.Build()
	if err != nil {
		common.Logger.Error("init logger error: ", zap.String("err", err.Error()))
	} else {
		common.Logger.Info("log init")
	}
	common.Logger.Sync()
}
func initAdvancedLogger() {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	topicDebugging := zapcore.AddSync(ioutil.Discard)
	topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core,it's easy to construct a Logger.
	common.Logger = zap.New(core)
}
