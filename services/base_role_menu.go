package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
)

//RoleMenuCreate menu create
func RoleMenuCreate(rm *models.RoleMenu) (*models.RoleMenu, error) {
	var (
		err error
	)
	rm, err = rm.Create(common.DB)
	return rm, err
}

//RoleMenuUpdate menu create
func RoleMenuUpdate(rm *models.RoleMenu) (*models.RoleMenu, error) {
	var (
		err error
	)
	rm, err = rm.Update(common.DB)
	return rm, err
}
func GetRoleMenuByID(id int) (*models.RoleMenu, error) {
	var (
		rm  models.RoleMenu
		err error
	)
	err = common.DB.First(&rm, "id = ?", id).Related(&rm.ResourceApi).Related(&rm.Menu).Error
	return &rm, err
}
