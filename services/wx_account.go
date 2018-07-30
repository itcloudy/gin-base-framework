package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
)

//GetOpenId  by id
func GetOpenId(openId string) (*models.WxAppId, error) {

	var (
		model models.WxAppId
		err   error
	)
	db := common.DB
	err = db.First(&model, "open_id = ?", openId).Error
	return &model, err
}

//OpenIdCreate   create  WxAppId
func OpenIdCreate(model *models.WxAppId) (*models.WxAppId, error) {

	var (
		validate *validator.Validate
		err      error
	)
	validate = validator.New()
	err = validate.Struct(model)
	if err != nil {
		return nil, err
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}

	}()
	model, err = model.Create(tx)

	return model, err
}
