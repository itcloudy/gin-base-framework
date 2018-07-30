package common

const (
	//request success
	SUCCESSED = 200
	//token error
	TOKENERR = 401
	//forbidden have no right for action
	FORBIDDEN = 403
	//request failed
	FAILED = 500
	//用户名或密码错误
	USERNAME_OR_PASSWORD_ERR = 101
	//数据已存在
	RECORD_EXISTED = 102
	//数据不存在
	RECORD_NOT_EXISTED = 104
	//无效的请求参数
	INVALID_PARAMETES = 106
	//数据删除失败
	DELETE_FAILED = 108
	// user menus get error
	MENU_GET_ERR = 109
	//文件上传获得文件失败
	UPLOAD_FILE_NO_EXSIT = 110
	//文件上传创建失败
	UPLOAD_FILE_CREATE_ERR = 111
	//获取文件资源失败
	UPLOAD_FILE_RESROUCE_ERR = 112
	//支付参数不对
	PAY_PARAMS_ERR = 113
	//支付随机数生成失败
	PAY_RAND_PARAM_ERR = 114
	// 支付请求失败
	PAY_REQUEST_ERR = 115
	// 支付请求失败
	PAY_ERR = 116
	//支付响应解析失败
	PAY_RESPONSE_UNMARSHAL_ERR = 117
	//支付失败
	PAY_RESULT_ERR = 118
)
const (
	// login user name
	LOGIN_USER_NAME = "LOGIN_USER_NAME"
	// login user id
	LOGIN_USER_ID = "LOGIN_USER_ID"
	// login user roles []string
	LOGIN_USER_ROLES = "LOGIN_USER_ROLES"
	//login user is admin
	LOGIN_IS_ADMIN = "LOGIN_IS_ADMIN"
	// token is valid
	TOKEN_VALID = "TOKEN_VALID"

)
const (
	//default page
	DEFAULT_PAGE = 1
	//default page limit
	DEFAULT_LIMIT = 20
	//default dorder
	DEFAULT_ORDER = "id desc"
)


const UPLOAD_FILE_URL = "/upload_files/"
const SYSTEM_STATIC_FILE_URL = "/system_statics/"
const TIME_FORMAT = "2006-01-02T15:04:05Z"
const TIME_FORMAT_ORDER = "20060102150405"
