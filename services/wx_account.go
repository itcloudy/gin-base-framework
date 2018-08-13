package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
)

//GetOpenId  by id
func GetOpenId(openId string) (*models.WxAppId, error, int) {
	var (
		model models.WxAppId
		code  int
		err   error
	)
	code = common.SUCCESSED
	db := common.DB
	err = db.Where("open_id = ?", openId).First(&model).Error
	if err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &model, err, code
}

//OpenIdCreate   create  WxAppId
func OpenIdCreate(model *models.WxAppId) (reModel *models.WxAppId, err error, code int) {

	var (
		validate *validator.Validate
	)
	validate = validator.New()
	err = validate.Struct(model)
	if err != nil {
		code = common.DATA_VALIDATE_ERR
		return
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}

	}()
	reModel, err = model.Create(tx)
	if err != nil {
		code = common.DB_INSERT_ERR
		return
	}
	code = common.SUCCESSED
	return
}
