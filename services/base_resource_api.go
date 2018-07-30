package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
)

//ResourceApiCreate menu create
func ResourceApiCreate(ra *models.ResourceApi) (*models.ResourceApi, error) {
	var (
		err error
	)

	ra, err = ra.Create(common.DB)
	return ra, err
}
func ResourceApiUpdate(ra *models.ResourceApi) (*models.ResourceApi, error) {
	var (
		err error
	)

	ra, err = ra.Update(common.DB)
	return ra, err
}
func GetResourceApi(method, address string, menuId int) (*models.ResourceApi, error) {
	var (
		rap models.ResourceApi
		err error
	)
	db := common.DB
	err = db.Where("menu_id = ? AND method = ? AND address = ?", menuId, method, address).Find(&rap).Error
	return &rap, err
}
