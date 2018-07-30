package common

import (
	"time"
)


type initModel struct {
	Menu         bool `yaml:"menu"`          // init menu information
	Role         bool `yaml:"role"`          // init role information
	Banner       bool `yaml:"banner"`        // init banner information
	CompanyLevel bool `yaml:"company_level"` // init company level information
}

//serverModel get server information from config.yml
type serverModel struct {
	Mode                 string        `yaml:"mode"`                    // run mode
	Host                 string        `yaml:"host"`                    // server host
	Port                 string        `yaml:"port"`                    // server port
	EnableHttps          bool          `yaml:"enable_https"`            // enable https
	CertFile             string        `yaml:"cert_file"`               // cert file path
	KeyFile              string        `yaml:"key_file"`                // key file path
	JwtPubKeyPath        string        `yaml:"jwt_public_key_path"`     // jwt public key path
	JwtPriKeyPath        string        `yaml:"jwt_private_key_path"`    // jwt private key path
	TokenExpireSecond    time.Duration `yaml:"token_expire_second"`     // token expire second
	SystemStaticFilePath string        `yaml:"system_static_file_path"` // system static file path
}

//databaseModel get database information from config.yml
type databaseModel struct {
	DBType  string `yaml:"type"`     // db type
	Connect string `yaml:"connect"`  // db connect information
	MaxIdle int    `yaml:"max_idle"` // db max idle connections
	MaxOpen int    `yaml:"max_open"` // db max open connections
}

//redisModel get redis information from config.yml
type redisModel struct {
	Address  string `yaml:"address"`
	Password string `yaml:"password"`
	DbName   int    `yaml:"db_name"`
}

//adminUserModel get admin user information from config.yml
type adminUserModel struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Email    string `yaml:"email"`
}

//weChatModel get admin user information from config.yml
type weChatModel struct {
	AppID      string `yaml:"app_id"`
	Secret     string `yaml:"secret"`
	OpenIdUrl  string `yaml:"openid_url"`
	PayAppID   string `yaml:"pay_app_id"`
	MchID      string `yaml:"mch_id"`
	NotifyUrl  string `yaml:"notify_url"`
	PayUrl     string `yaml:"pay_url"`
	PaySignKey string `yaml:"pay_sign_key"`
}

//fileUploadModel get file upload information from config.yml
type fileUploadModel struct {
	Storage string `yaml:"storage"`
	Path    string `yaml:"path"`
}

//qiNiuModel get qi niu cloud storage information from config.yml
type qiNiuModel struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Bucket    string `yaml:"bucket"`
	Domain    string `yaml:"domain"`
}

type messageQueueModel struct {
	Enable bool   `yaml:"enable"`
	Type   string `yaml:"type"`
	Redis  struct {
	} `yaml:"redis"`
	Kafka struct {
	} `yaml:"kafka"`
	Amqp struct {
		Url string `yaml:"url"`
	} `yaml:"amqp"`
}

type logModel struct {
	Level      string `yaml:"level"`
	ConfigFile string `yaml:"config_file"`
}
type configModel struct {
	Server              *serverModel                      `yaml:"server"`
	Database            *databaseModel                    `yaml:"database"`
	Redis               *redisModel                       `yaml:"redis"`
	AdminUser           *adminUserModel                   `yaml:"admin_user"`
	WeChat              *weChatModel                      `yaml:"we_chat"`
	FileUpload          *fileUploadModel                  `yaml:"file_upload"`
	QiNiu               *qiNiuModel                       `yaml:"qi_niu"`
	SendMessageQueue    *messageQueueModel                `yaml:"send_message_queue"`
	ReceiveMessageQueue *messageQueueModel                `yaml:"receive_message_queue"`
	Session             *serverModel                      `yaml:"session"`
	Log                 *logModel                         `yaml:"log"`
	Init                *initModel                        `yaml:"init"`
}
