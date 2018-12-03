package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/itcloudy/gin-base-framework/models"
	"github.com/itcloudy/gin-base-framework/services"
	"github.com/pkg/errors"
	"strconv"
)

// @tags  用户
// @Description 个人信息获得
// @Summary 个人信息获得
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":[],"openid":"","active":true,"is_admin":true},"message":"success"}}"
// @Router /auth/self [get]
func SelfInfo(c *gin.Context) {

	var (
		user *models.User
		err  error
		code int
	)
	id := c.GetInt(common.LOGIN_USER_ID)
	user, err, code = services.GetUserById(id)
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, user, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}

}

// @tags  用户
// @Description 个人信息修改
// @Summary 个人信息修改
// @Accept  multipart/form-data
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":[],"openid":"","active":true,"is_admin":true},"message":"success"}}"
// @Router /auth/self [put]
func SelfInfoUpdate(c *gin.Context) {

	var (
		user *models.User
		err  error
		code int
	)
	id := c.GetInt(common.LOGIN_USER_ID)
	user.ID = id

	if name := c.Query("name"); name != "" {
		user.Name = name
	}
	if email := c.Query("email"); email != "" {
		user.Email = email
	}
	if alias := c.Query("alias"); alias != "" {
		user.Alias = alias
	}
	user, err, code = services.UserUpdate(user, c.GetBool(common.LOGIN_IS_ADMIN))
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, user, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}

}

// @tags  用户
// @Description 用户列表，只有管理员才可调用
// @Summary 用户列表
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":[{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":null,"openid":"","active":true,"is_admin":true},{"id":2,"name":"admin","alias":"","email":"","password":"","roles":null,"openid":"admin","active":true,"is_admin":false}],"message":"success"}"
// @Router /auth/users [get]
func UserGetAll(c *gin.Context) {
	var (
		err        error
		userList   []*models.User
		whereArgs  []interface{}
		whereQuery string
		code       int
		total      int
	)
	if !c.GetBool(common.LOGIN_IS_ADMIN) {
		common.GenResponse(c, code, nil, "you have no right to do this")
		return
	}
	page := common.String2Int(c.Query("page"), common.DEFAULT_PAGE)
	limit := common.String2Int(c.Query("limit"), common.DEFAULT_LIMIT)
	condition := make(map[string]interface{})
	condition["page"] = page
	condition["limit"] = limit
	name := c.Query("name")
	if name != "" {
		whereArgs = append(whereArgs, name)
		/*if whereQuery ==""{
			whereQuery = " name = ? "
		}else{
			whereQuery += "AND name = ? "
		}*/
	}
	email := c.Query("email")
	if email != "" {
		whereArgs = append(whereArgs, email)
		if whereQuery == "" {
			whereQuery = " email = ? "
		} else {
			whereQuery += "AND email = ? "
		}
	}
	data := make(map[string]interface{})
	userList, total, err, code = services.GetAllUser(whereQuery, whereArgs, "id desc", page, limit)
	if err == nil {
		data["data_list"] = userList
		data["total"] = total

		common.GenResponse(c, common.SUCCESSED, userList, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}

}

// @tags  用户
// @Description 用户获得,只有管理员才可调用
// @Summary 用户获得
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path int true "用户ID"
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":[],"openid":"","active":true,"is_admin":true},"message":"success"}}"
// @Router /auth/user/{id} [get]
func UserGet(c *gin.Context) {
	//TODO department permission control
	var (
		code int
	)
	if !c.GetBool(common.LOGIN_IS_ADMIN) {
		code = common.FORBIDDEN
		common.GenResponse(c, code, nil, "you have no right to do this")
		return
	}
	var (
		user *models.User
		err  error
		id   int
	)
	id, err = strconv.Atoi(c.Param("id"))
	if err == nil {
		user, err, code = services.GetUserById(id)
	} else {
		err = errors.New("param id empty or not a number")
	}
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, user, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}

}

// @tags  用户
// @Description 用户注册
// @Summary 用户注册
// @Accept  multipart/form-data
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":[],"openid":"","active":true,"is_admin":true},"message":"success"}}"
// @Router /auth/user/ [post]
func UserRegister(c *gin.Context) {
	var (
		user   models.User
		reUser *models.User
		err    error
		code   int
	)
	err = c.ShouldBindWith(&user, binding.JSON)
	if err != nil {
		code = common.BINDING_JSON_ERR
		common.GenResponse(c, code, nil, err.Error())
		return
	}
	//user register can't add any role
	user.RoleList = []int{}
	reUser, err, code = services.UserCreate(&user)
	if err == nil {
		user.Password = ""
		common.GenResponse(c, common.SUCCESSED, reUser, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}

// @tags  用户
// @Description 用户信息修改,只有管理员才可调用，不修改头像
// @Summary 用户信息修改
// @Accept  json
// @Produce  json
// @Param id path int true "用户信息ID"
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":{"id":1,"name":"admin","alias":"","email":"admin@block.vc","password":"","roles":[],"openid":"","active":true,"is_admin":true},"message":"success"}}"
// @Router /auth/user/{id} [put]
func UserUpdate(c *gin.Context) {
	var (
		user   models.User
		reUser *models.User
		err    error
		code   int
	)

	id, err := strconv.Atoi(c.Param("id"))

	//only admin or user self can update
	if !c.GetBool(common.LOGIN_IS_ADMIN) {
		code = common.FORBIDDEN
		common.GenResponse(c, code, nil, "you have no right to do this")
		return
	} else {
		if c.GetInt(common.LOGIN_USER_ID) != id {
			common.GenResponse(c, code, nil, "you have no right to do this")
			return
		}
	}

	if err = c.ShouldBindWith(&user, binding.JSON); err != nil {
		common.GenResponse(c, code, nil, err.Error())
	}
	user.ID = id
	reUser, err, code = services.UserUpdate(&user, c.GetBool(common.LOGIN_IS_ADMIN))
	if err == nil {
		user.Password = ""
		common.GenResponse(c, common.SUCCESSED, reUser, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}
