package common

func GetResponseMessage(code int) (message string) {
	switch code {
	case SUCCESSED:
		message = "操作成功"
	case DATA_VALIDATE_ERR:
		message = "数据验证失败"
	case DB_INSERT_ERR:
		message = "记录创建失败"
	case DB_UPDATE_ERR:
		message = "记录更新失败"
	case DB_DELETE_ERR:
		message = "记录删除失败"
	case BINDING_JSON_ERR:
		message = "数据绑定失败"
	case DB_RECORD_NOT_FOUND:
		message = "记录查询失败"
	case REQUEST_DATA_EMPITY:
		message = "请求参数为空"
	case SYSTEM_HAS_NO_MENUS:
		message = "系统没有菜单"
	case ROLE_HAS_NO_MENUS:
		message = "角色没有菜单"
	case ROLE_MENUS_TREE_ERR:
		message = "菜单树不存在"
	case SYSTEM_HAS_NO_USER:
		message = "系统用户不存在"
	case NAME_PASSWORD_INVALID:
		message = "用户名或密码无效"
	case MOBILE_REPEAT:
		message = "手机号码重复"
	case OPEN_ID_IS_EMPITY:
		message = "openid 为空"
		message = "需要登录"
	case TOKENERR:
		message = "Token错误"
	case FORBIDDEN:
		message = "无权操作"
	case FAILED:
		message = "操作失败"
	case USERNAME_OR_PASSWORD_ERR:
		message = "用户名或密码错误"
	case RECORD_EXISTED:
		message = "数据已存在"
	case MENU_GET_ERR:
		message = "用户菜单获取失败"
	case UPLOAD_FILE_NO_EXSIT:
		message = "文件上传获得文件失败"
	case UPLOAD_FILE_CREATE_ERR:
		message = "文件上传创建失败"
	case INVALID_PARAMETES:
		message = "无效的请求参数"
	case UPLOAD_FILE_RESROUCE_ERR:
		message = "获取文件资源失败"
	case PAY_PARAMS_ERR:
		message = "支付参数不对"
	case PAY_RAND_PARAM_ERR:
		message = "支付随机数生成失败"
	case PAY_REQUEST_ERR:
		message = "支付请求失败"
	case PAY_ERR:
		message = "支付请求失败"
	case PAY_RESPONSE_UNMARSHAL_ERR:
		message = "支付响应解析失败"
	case PAY_RESULT_ERR:
		message = "支付失败"
	case PAY_RESPONSE_ERR:
		message = "支付响应失败"
	case PAY_REQUEST_POST_ERR:
		message = "支付请求失败" //
	case PASSWORD_INVALID:
		message = "密码错误"

	default:
		message = "未知错误"

	}
	return
}
