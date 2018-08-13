package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
)

//RoleMenuCreate menu create
func RoleMenuCreate(rm *models.RoleMenu) (*models.RoleMenu, error, int) {
	var (
		err  error
		code int
	)
	code = common.SUCCESSED

	if rm, err = rm.Create(common.DB); err != nil {
		code = common.DB_INSERT_ERR
	}
	return rm, err, code
}

//RoleMenuUpdate menu create
func RoleMenuUpdate(rm *models.RoleMenu) (reModel *models.RoleMenu, err error, code int) {

	code = common.SUCCESSED

	if rm, err = rm.Update(common.DB); err != nil {
		code = common.DB_UPDATE_ERR
	}
	return
}
func GetRoleMenuByID(id int) (*models.RoleMenu, error, int) {
	var (
		rm   models.RoleMenu
		err  error
		code int
	)
	code = common.SUCCESSED

	if err = common.DB.First(&rm, "id = ?", id).Related(&rm.Menu).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &rm, err, code
}
