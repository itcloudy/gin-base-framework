package common

const (
	// 数据验证失败
	DATA_VALIDATE_ERR = 2000
	// 记录创建失败
	DB_INSERT_ERR = 2001
	// 记录更新失败
	DB_UPDATE_ERR = 2002
	// 记录删除失败
	DB_DELETE_ERR = 2003
	// 数据绑定失败
	BINDING_JSON_ERR = 2004
	// 记录查询失败
	DB_RECORD_NOT_FOUND = 2005
	// 请求参数为空
	REQUEST_DATA_EMPITY = 2006
	// 系统没有菜单
	SYSTEM_HAS_NO_MENUS = 2007
	// 角色没有菜单
	ROLE_HAS_NO_MENUS = 2008
	// 菜单树不存在
	ROLE_MENUS_TREE_ERR = 2009
	// 系统用户不存在
	SYSTEM_HAS_NO_USER = 2010
	// 用户名或密码无效
	NAME_PASSWORD_INVALID = 2011
	// 手机号码重复
	MOBILE_REPEAT = 2012
	// openid 为空
	OPEN_ID_IS_EMPITY = 2013
	//itcloudy登录失败
	// 密码错误
	PASSWORD_INVALID = 2022
	// 请求成功
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

	//无效的请求参数
	INVALID_PARAMETES = 106

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
	// 支付响应失败
	PAY_RESPONSE_ERR = 119
	// 支付请求失败
	PAY_REQUEST_POST_ERR = 120
	// openID获取失败
	GET_OPEN_ID_FAILED = 121
)
const (
	// Accept-Language
	ACCEPT_LANGUAGE = "ACCEPT_LANGUAGE"
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
	// login user company id
	LOGIN_COMPANY_ID = "LOGIN_COMPANY_ID"
	// login user company NO.
	LOGIN_COMPANY_REQUEST_NO = "LOGIN_COMPANY_REQUEST_NO"
)
const (
	//default page
	DEFAULT_PAGE = 1
	//default page limit
	DEFAULT_LIMIT = 15
	//default dorder
	DEFAULT_ORDER = "id desc"
)


const UPLOAD_FILE_URL = "/upload_files/"
const SYSTEM_STATIC_FILE_URL = "/system_statics/"
const UPLOAD_FILE_MIME = "jpg, png"
const TIME_FORMAT = "2006-01-02T15:04:05Z"
const TIME_FORMAT_ORDER = "20060102150405"
