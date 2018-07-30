package services

import (
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/daemons"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
)

//RoleCreate role create
func RoleCreate(role *models.Role) (*models.Role, error) {

	var (
		validate           *validator.Validate
		err                error
		txErr              error
		policyActions      []common.PolicyAction
		groupPolicyActions []common.GroupPolicyAction
	)
	validate = validator.New()
	err = validate.Struct(role)
	if err != nil {
		return nil, err
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			txErr = tx.Commit().Error
		} else {
			txErr = tx.Rollback().Error
		}
		if txErr == nil && err == nil {
			pType := fmt.Sprintf("role_%s", role.ID)
			daemons.RoleMenusEnforcerDaemon(pType, policyActions, role.ID)
			daemons.UserOrRoleEnforcerDaemon(groupPolicyActions)
		}
		if err != nil {
			err = txErr
		}

	}()
	role, err = role.Create(tx)

	if err == nil {
		for _, ma := range role.MenuApis {
			for _, apiId := range ma.ApiIds {

				var rm models.RoleMenu
				rm.MenuID = ma.MenuID
				rm.RoleID = role.ID
				rm.RoleID = role.ID
				rm.ResourceApiID = apiId
				rm.Create(tx)
			}
		}
		for _, ro := range role.Inherits {
			var gpa common.GroupPolicyAction
			gpa.Action = "add"
			gpa.Role = fmt.Sprintf("role_%d", role.ID)
			gpa.UserOrRole = fmt.Sprintf("role_%d", ro.ID)
			groupPolicyActions = append(groupPolicyActions, gpa)

		}
	}

	return role, err

}

//RoleUpdate menu create
func RoleUpdate(role *models.Role) (*models.Role, error) {
	var (
		err                error
		txErr              error
		oldRmMenus         []models.RoleMenu
		policyActions      []common.PolicyAction
		groupPolicyActions []common.GroupPolicyAction
	)

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			txErr = tx.Commit().Error
		} else {
			txErr = tx.Rollback().Error
		}
		if txErr == nil && err == nil {
			pType := fmt.Sprintf("role_%s", role.ID)
			daemons.RoleMenusEnforcerDaemon(pType, policyActions, role.ID)
			daemons.UserOrRoleEnforcerDaemon(groupPolicyActions)
		}
		if err != nil {
			err = txErr
		}

	}()
	// 默认用户的名称不可修改，id为1的
	if role.ID == 1 {
		role.Name = ""
	}
	role, err = role.Update(tx)
	if err == nil {
		//delete all role menus
		tx.Where("role_id = ?", role.ID).Find(&oldRmMenus)
		tx.Where("id = ?", role.ID).First(&role)
		for _, ro := range role.Inherits {
			var gpa common.GroupPolicyAction
			gpa.Action = "delete"
			gpa.Role = fmt.Sprintf("role_%d", role.ID)
			gpa.UserOrRole = fmt.Sprintf("role_%d", ro.ID)
			groupPolicyActions = append(groupPolicyActions, gpa)
		}
		for _, rm := range oldRmMenus {
			if r, e := GetRoleMenuByID(rm.ID); e == nil {
				var pA common.PolicyAction
				pA.Method = r.ResourceApi.Method
				pA.Address = r.ResourceApi.Address
				pA.PType = fmt.Sprintf("role_%d", role.ID)
				policyActions = append(policyActions, pA)
			}

		}
		tx.Where("role_id = ?", role.ID).Delete(models.RoleMenu{})
		// create role menus
		for _, ma := range role.MenuApis {
			for _, apiId := range ma.ApiIds {

				var rm models.RoleMenu
				rm.MenuID = ma.MenuID
				rm.RoleID = role.ID
				rm.RoleID = role.ID
				rm.ResourceApiID = apiId
				rm.Create(tx)
			}
		}
		for _, ro := range role.Inherits {
			var gpa common.GroupPolicyAction
			gpa.Action = "add"
			gpa.Role = fmt.Sprintf("role_%d", role.ID)
			gpa.UserOrRole = fmt.Sprintf("role_%d", ro.ID)
			groupPolicyActions = append(groupPolicyActions, gpa)

		}
	}

	return role, err

}

//GetRoleById get menu by id
func GetRoleById(id int) (*models.Role, error) {

	var (
		role models.Role
		err  error
	)
	db := common.DB
	err = db.First(&role, "id = ?", id).Related(&role.RoleMenus, "RoleMenus").Error

	for _, m := range role.RoleMenus {
		if m.ResourceApiID > 0 {
			db.First(&m.ResourceApi)
		}
		if m.MenuID > 0 {
			db.First(&m.Menu)
		}
		if m.RoleID > 0 {
			db.First(&m.Role)
		}
	}

	return &role, err
}

func GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := common.DB.First(&role, "name = ?", name).Related(&role.RoleMenus, "RoleMenus").Error
	return &role, err
}

//DeleteRoleById get category by id
func DeleteRoleById(id int) (*models.Role, error) {

	var (
		model models.Role
		err   error
	)
	model.ID = id
	err = common.DB.Delete(&model).Error
	return nil, err
}

//GetAllRole get all banner_group
func GetAllRole(where map[string]interface{}, page int, limit int, order string) ([]*models.Role, error) {
	var (
		err      error
		roleList []*models.Role
	)
	db := common.DB
	err = db.Where(where).Offset((page - 1) * limit).Limit(limit).Find(&roleList).Order(order).Error
	return roleList, err
}
