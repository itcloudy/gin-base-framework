package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

//SystemApiCreate   create  collection
func SystemApiCreate(model *models.SystemApi) (*models.SystemApi, error, int) {

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

//GetAllSystemApiFromDB get all
func GetAllSystemApiFromDB() []*models.SystemApi {
	var (
		api []*models.SystemApi
	)
	common.DB.Find(&api)
	return api
}

func GetSystemAPIByThreeParams(name, method, address string) (*models.SystemApi, error, int) {
	var (
		model models.SystemApi
		err   error
		code  int
	)
	code = common.SUCCESSED
	method = strings.ToUpper(method)
	if err = common.DB.Where("name = ? AND method = ? AND address = ?", name, method, address).First(&model).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}

	return &model, err, code

}
