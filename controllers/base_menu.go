package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"strconv"
	"strings"
)

// @tags  菜单
// @Description 菜单获得
// @Summary 菜单获得
// @Accept  json
// @Produce  json
// @Param id path int true "菜单获得ID"
// @Param Authorization header string true "Token"
// @Router /auth/menu/{id} [get] "{"code":200,"data":[{"id":1,"parent_id":0,"name":"系统设置","route":"route-system","component":"SystemComponent","icon":"system-icon","sequence":1,"children":[{"id":2,"parent_id":1,"name":"系统用户","route":"/system/user","icon":"system-user","sequence":1}]},{"id":3,"parent_id":0,"name":"分类模块","route":"route_modyle","component":"Module","icon":"module","sequence":2,"children":[{"id":4,"parent_id":3,"name":"模块列表","route":"/auth/category","icon":"category","sequence":1}]},{"id":5,"parent_id":0,"name":"公司管理","route":"route_company","component":"Company","icon":"company","sequence":2,"children":[{"id":6,"parent_id":5,"name":"公司分类","route":"route_company_1","component":"Company","icon":"company","sequence":2,"children":[{"id":7,"parent_id":6,"name":"一级公司","route":"route_company_2","component":"Company","icon":"company","sequence":2},{"id":8,"parent_id":6,"name":"二级公司","route":"route_company_3","component":"Company","icon":"company","sequence":2}]}]}],"message":"success"}"
func MenuGet(c *gin.Context) {
	var (
		menu *models.Menu
		err  error
		code int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		menu, err, code = services.GetMenuById(id)
	} else {
		code = common.REQUEST_DATA_EMPITY
		err = errors.New("param id empty or not a number")
	}
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, menu, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}

// @tags  菜单
// @Description 菜单修改
// @Summary 菜单修改
// @Accept  json
// @Produce  json
// @Param id path int true "轮播图组ID"
// @Param Authorization header string true "Token"
// @Param active query bool true "是否有效"
// @Param start_date query time true "开始时间"
// @Param end_date query time true "结束时间"
// @Param name query string true "轮播组名称"
// @Param display_time query int true "轮播时间间隔(s)"
// @Success 200 {string} json  "{"code":200,"data":{"id":1,"parent_id":0,"name":"测试修改","route":"route-system","component":"SystemComponent","icon":"system-icon","sequence":1},"message":"success"}"
// @Router /auth/menu/{id} [put]
func MenuUpdate(c *gin.Context) {
	var (
		menu   models.Menu
		reMenu *models.Menu
		err    error
		code   int
	)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		code = common.REQUEST_DATA_EMPITY
		common.GenResponse(c, code, nil, "param id empty or not a number")
		return
	}
	err = c.ShouldBindWith(&menu, binding.JSON)
	if err != nil {
		common.GenResponse(c, code, nil, err.Error())
		return
	} else {
		menu.ID = id
		reMenu, err, code = services.MenuUpdate(&menu)
		if err == nil {
			common.GenResponse(c, common.SUCCESSED, reMenu, "success")
		} else {
			common.GenResponse(c, code, nil, err.Error())
		}
	}
}

// @tags  菜单
// @Description 用户菜单树获得
// @Summary 用户菜单树获得
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":[{"id":1,"parent_id":0,"name":"系统设置","route":"route-system","component":"SystemComponent","icon":"system-icon","sequence":1,"children":[{"id":2,"parent_id":1,"name":"系统用户","route":"/system/user","icon":"system-user","sequence":1}]},{"id":3,"parent_id":0,"name":"分类模块","route":"route_modyle","component":"Module","icon":"module","sequence":2,"children":[{"id":4,"parent_id":3,"name":"模块列表","route":"/auth/category","icon":"category","sequence":1}]},{"id":5,"parent_id":0,"name":"公司管理","route":"route_company","component":"Company","icon":"company","sequence":2,"children":[{"id":6,"parent_id":5,"name":"公司分类","route":"route_company_1","component":"Company","icon":"company","sequence":2,"children":[{"id":7,"parent_id":6,"name":"一级公司","route":"route_company_2","component":"Company","icon":"company","sequence":2},{"id":8,"parent_id":6,"name":"二级公司","route":"route_company_3","component":"Company","icon":"company","sequence":2}]}]}],"message":"success"}}"
// @Router /auth/usermenu [get]
func UserMenu(c *gin.Context) {
	var (
		menus   []*models.Menu
		err     error
		roleIds []int
		code    int
	)

	roleList := c.GetStringSlice(common.LOGIN_USER_ROLES)
	if len(roleList) > 0 {
		for _, r := range roleList {
			m := strings.Split(r, "_")
			if roleId, e := strconv.Atoi(m[1]); e == nil {
				roleIds = append(roleIds, roleId)
			}
		}
	}
	menus, err, code = services.GetMenuByUserRoleIds(roleIds, false)

	if err == nil {
		if len(menus) == 0 {
			menus = []*models.Menu{}
		}
		common.GenResponse(c, common.SUCCESSED, menus, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}

}

// @tags  菜单
// @Description 系统菜单树,系统管理员才可用
// @Summary 系统菜单树
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":[{"id":1,"parent_id":0,"name":"系统设置","route":"route-system","component":"SystemComponent","icon":"system-icon","sequence":1,"children":[{"id":2,"parent_id":1,"name":"系统用户","route":"/system/user","icon":"system-user","sequence":1}]},{"id":3,"parent_id":0,"name":"分类模块","route":"route_modyle","component":"Module","icon":"module","sequence":2,"children":[{"id":4,"parent_id":3,"name":"模块列表","route":"/auth/category","icon":"category","sequence":1}]},{"id":5,"parent_id":0,"name":"公司管理","route":"route_company","component":"Company","icon":"company","sequence":2,"children":[{"id":6,"parent_id":5,"name":"公司分类","route":"route_company_1","component":"Company","icon":"company","sequence":2,"children":[{"id":7,"parent_id":6,"name":"一级公司","route":"route_company_2","component":"Company","icon":"company","sequence":2},{"id":8,"parent_id":6,"name":"二级公司","route":"route_company_3","component":"Company","icon":"company","sequence":2}]}]}],"message":"success"}}"
// @Router /auth/menutree [get]
func MenuTree(c *gin.Context) {
	var (
		menus   []*models.Menu
		err     error
		roleIds []int
		code    int
	)
	if c.GetBool(common.LOGIN_IS_ADMIN) {
		menus, err, code = services.GetMenuByUserRoleIds(roleIds, true)
		if err == nil {
			common.GenResponse(c, common.SUCCESSED, menus, "success")
		} else {
			common.GenResponse(c, code, nil, err.Error())
		}
	} else {
		common.GenResponse(c, common.FORBIDDEN, nil, "you have no right to do for this action")
	}

}

// @tags  菜单
// @Description 用户菜单树,系统管理员才可用
// @Summary 用户菜单树
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param user_id path int true "用户ID"
// @Success 200 {string} json "{"code":200,"data":[{"id":1,"parent_id":0,"name":"系统设置","route":"route-system","component":"SystemComponent","icon":"system-icon","sequence":1,"children":[{"id":2,"parent_id":1,"name":"系统用户","route":"/system/user","icon":"system-user","sequence":1}]},{"id":3,"parent_id":0,"name":"分类模块","route":"route_modyle","component":"Module","icon":"module","sequence":2,"children":[{"id":4,"parent_id":3,"name":"模块列表","route":"/auth/category","icon":"category","sequence":1}]},{"id":5,"parent_id":0,"name":"公司管理","route":"route_company","component":"Company","icon":"company","sequence":2,"children":[{"id":6,"parent_id":5,"name":"公司分类","route":"route_company_1","component":"Company","icon":"company","sequence":2,"children":[{"id":7,"parent_id":6,"name":"一级公司","route":"route_company_2","component":"Company","icon":"company","sequence":2},{"id":8,"parent_id":6,"name":"二级公司","route":"route_company_3","component":"Company","icon":"company","sequence":2}]}]}],"message":"success"}}"
// @Router /auth/usermenutree/{user_id} [get]
func UserMenuTree(c *gin.Context) {
	var (
		menus []*models.Menu
		err   error
		code  int
	)
	if c.GetBool(common.LOGIN_IS_ADMIN) {
		if userId, err := strconv.Atoi(c.Param("user_id")); err == nil {
			menus, err, code = services.GetMenuByUserID(userId)
		}

		if err == nil {
			common.GenResponse(c, common.SUCCESSED, menus, "success")
		} else {
			common.GenResponse(c, code, nil, err.Error())
		}
	} else {
		common.GenResponse(c, common.FORBIDDEN, nil, "you have no right for this action")
	}

}

// @tags  菜单
// @Description 菜单删除
// @Summary 菜单删除
// @Accept  json
// @Produce  json
// @Param id path int true "菜单ID"
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":null,"message":"success"}"
// @Router /auth/menu/{id} [delete]
func MenuDelete(c *gin.Context) {

	var (
		err  error
		code int
	)
	if !c.GetBool(common.LOGIN_IS_ADMIN) {
		common.GenResponse(c, common.FORBIDDEN, nil, "you have no right for this action")

	}
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		_, err, code = services.DeleteMenupById(id)
	} else {
		code = common.REQUEST_DATA_EMPITY
		err = errors.New("param id empty or not a number")
	}
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, nil, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}
