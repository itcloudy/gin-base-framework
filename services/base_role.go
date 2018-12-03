package services

import (
	"fmt"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/itcloudy/gin-base-framework/models"
	"github.com/itcloudy/validator"
)

//RoleCreate role create
func RoleCreate(role *models.Role) (*models.Role, error, int) {

	var (
		validate *validator.Validate
		err      error
		code     int
	)
	code = common.SUCCESSED
	validate = validator.New()
	err = validate.Struct(role)
	if err != nil {
		code = common.DATA_VALIDATE_ERR
		return nil, err, code
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()
	if role, err = role.Create(tx); err != nil {
		code = common.DB_INSERT_ERR
	}

	return role, err, code

}

//RoleUpdate menu create
func RoleUpdate(role *models.Role) (*models.Role, error, int) {
	var (
		err                error
		oldRmMenus         []models.RoleMenu
		policyActions      []common.PolicyAction
		groupPolicyActions []common.GroupPolicyAction
		code               int
	)
	code = common.SUCCESSED

	validate := validator.New()
	err = validate.Struct(role)
	if err != nil {
		code = common.DATA_VALIDATE_ERR
		return nil, err, code
	}
	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}

	}()
	// 默认用户的名称不可修改，id为1的
	if role.ID == 1 {
		role.Name = ""
	}
	if role, err = role.Update(tx); err != nil {
		code = common.DB_UPDATE_ERR
		return nil, err, code
	}
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
			if _, e, _ := GetRoleMenuByID(rm.ID); e == nil {
				var pA common.PolicyAction
				pA.PType = fmt.Sprintf("role_%d", role.ID)
				policyActions = append(policyActions, pA)
			}

		}
		tx.Where("role_id = ?", role.ID).Delete(models.RoleMenu{})

		for _, ro := range role.Inherits {
			var gpa common.GroupPolicyAction
			gpa.Action = "add"
			gpa.Role = fmt.Sprintf("role_%d", role.ID)
			gpa.UserOrRole = fmt.Sprintf("role_%d", ro.ID)
			groupPolicyActions = append(groupPolicyActions, gpa)

		}
	}

	return role, err, code

}

//GetRoleById get menu by id
func GetRoleById(id int) (*models.Role, error, int) {

	var (
		role models.Role
		err  error
		code int
	)
	code = common.SUCCESSED
	db := common.DB
	if err = db.First(&role, "id = ?", id).Related(&role.RoleMenus, "RoleMenus").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	for _, m := range role.RoleMenus {
		if m.MenuID > 0 {
			db.First(&m.Menu)
		}
		if m.RoleID > 0 {
			db.First(&m.Role)
		}
	}

	return &role, err, code
}

//GetRoleByCode get menu by code
func GetRoleByCode(uniqueCode string) (*models.Role, error, int) {

	var (
		role models.Role
		err  error
		code int
	)
	db := common.DB
	if err = db.First(&role, "code = ?", uniqueCode).Related(&role.RoleMenus, "RoleMenus").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}

	for _, m := range role.RoleMenus {

		if m.MenuID > 0 {
			db.First(&m.Menu)
		}
		if m.RoleID > 0 {
			db.First(&m.Role)
		}
	}

	return &role, err, code
}

func GetRoleByName(name string) (*models.Role, error, int) {
	var (
		role models.Role
		err  error
		code int
	)
	code = common.SUCCESSED
	if err = common.DB.First(&role, "name = ?", name).Related(&role.RoleMenus, "RoleMenus").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &role, err, code
}

//DeleteRoleById get role by id
func DeleteRoleById(id int) (*models.Role, error, int) {

	var (
		model models.Role
		err   error
		code  int
	)
	model.ID = id
	if err = common.DB.Delete(&model).Error; err != nil {
		code = common.DB_DELETE_ERR
	}
	return nil, err, code
}

//GetAllRole get all banner_group
func GetAllRole(where map[string]interface{}, page int, limit int, order string) ([]*models.Role, error, int) {
	var (
		err      error
		roleList []*models.Role
		code     int
	)
	db := common.DB
	if err = db.Where(where).Offset((page - 1) * limit).Limit(limit).Find(&roleList).Order(order).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return roleList, err, code
}

//GetAllRoleFromDB get all
func GetAllRoleFromDB() []*models.Role {
	var (
		roles []*models.Role
	)
	common.DB.Find(&roles)
	return roles
}
