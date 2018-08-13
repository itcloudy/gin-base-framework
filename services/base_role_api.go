package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
)

//RoleApiCreate   create  role api
func RoleApiCreate(model *models.RoleApi) (*models.RoleApi, error, int) {

	var (
		validate *validator.Validate
		err      error
		code     int
	)
	code = common.SUCCESSED
	validate = validator.New()
	err = validate.Struct(model)
	if err != nil {
		code = common.DATA_VALIDATE_ERR
		return nil, err, code
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}

	}()
	if model, err = model.Create(tx); err != nil {
		code = common.DB_INSERT_ERR
	}
	return model, err, code
}

func GetRoleAPIsByRoleId(roleId int) ([]*models.RoleApi, error, int) {
	var (
		roleAPis []*models.RoleApi
		err      error
		code     int
	)
	code = common.SUCCESSED
	if err = common.DB.Model(models.RoleApi{}).Preload("SystemApi").Where("role_id = ?", roleId).Find(&roleAPis).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return roleAPis, err, code

}
