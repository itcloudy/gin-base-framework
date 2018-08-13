package controllers

import (
	"errors"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strconv"
)

// @tags  角色
// @Description 角色创建
// @Summary 角色创建
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param name query string true "角色名称"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /auth/role [post]
func RoleCreate(c *gin.Context) {
	var (
		role   models.Role
		reRole *models.Role
		err    error
		code   int
	)
	err = c.ShouldBindWith(&role, binding.JSON)
	reRole, err, code = services.RoleCreate(&role)
	if err != nil {
		common.GenResponse(c, code, nil, err.Error())
	} else {
		common.GenResponse(c, common.SUCCESSED, reRole, "success")
	}
}

// @tags  角色
// @Description 角色更新
// @Summary 角色更新
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path int true "角色ID"
// @Param name query string true "角色名称"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /auth/role/{id} [put]
func RoleUpdate(c *gin.Context) {
	var (
		role   models.Role
		reRole *models.Role
		err    error
		code   int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = common.REQUEST_DATA_EMPITY
		common.GenResponse(c, code, nil, "param id empty or not a number")

		return
	}
	err = c.ShouldBindWith(&role, binding.JSON)
	if err != nil {
		common.GenResponse(c, code, nil, err.Error())

		return
	} else {

		role.ID = id
		reRole, err, code = services.RoleUpdate(&role)
		if err == nil {
			common.GenResponse(c, common.SUCCESSED, reRole, "success")
		} else {
			common.GenResponse(c, code, nil, err.Error())
		}
	}
}

// @tags  角色
// @Description 角色获得
// @Summary 角色获得
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path int true "角色ID"
// @Param name query string true "角色名称"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /auth/role/{id} [get]
func RoleGet(c *gin.Context) {

	var (
		role *models.Role
		err  error
		code int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		role, err, code = services.GetRoleById(id)
	} else {
		code = common.REQUEST_DATA_EMPITY
		err = errors.New("param id empty or not a number")
	}
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, role, "success")

	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}

// @tags  角色
// @Description 角色删除
// @Summary 角色列表
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param id path int true "角色ID"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /auth/role/{id} [delete]
func RoleDelete(c *gin.Context) {
	var (
		err  error
		code int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		_, err, code = services.DeleteRoleById(id)
	} else {
		err = errors.New("param id empty or not a number")
	}
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, nil, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}

// @tags  角色
// @Description 角色列表
// @Summary 角色列表
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param page query int false "页码"
// @Param limit query int false "每页数量"
// @Param order query string false "排序"
// @Success 200 {string} json "{"code":200,"data":{},"message":"ok"}"
// @Router /auth/roles [get]
func RoleAll(c *gin.Context) {
	var (
		page    int
		limit   int
		order   string
		err     error
		resInfo []*models.Role
		code    int
	)

	Where := make(map[string]interface{})
	page = common.String2Int(c.Query("page"), common.DEFAULT_PAGE)
	limit = common.String2Int(c.Query("limit"), common.DEFAULT_LIMIT)
	if order = c.Query("order"); order == "" {
		order = common.DEFAULT_ORDER
	}

	if c.Request.Form == nil {
		c.Request.ParseMultipartForm(32 << 20)
	}

	args := c.Request.Form

	for k, v := range args {
		if k == "page" || k == "limit" || k == "order" {
			continue
		}
		Where[k] = v[0]
	}

	if resInfo, err, code = services.GetAllRole(Where, page, limit, order); err != nil {
		common.GenResponse(c, code, err.Error(), "error")
	} else {
		common.GenResponse(c, common.SUCCESSED, resInfo, "success")
	}
}
