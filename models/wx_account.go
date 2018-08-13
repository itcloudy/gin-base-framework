package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type WxAppId struct {
	BaseModel
	OpenId string `gorm:"unique" json:"open_id" form:"open_id"` //微信id
}

func (model *WxAppId) Create(db *gorm.DB) (*WxAppId, error) {
	var wx WxAppId
	err := db.First(&wx, "open_id = ?", model.OpenId).Error
	if err == nil {
		err = errors.New("open_id is existed!")
		return nil, err
	}
	err = db.Create(model).Error
	db.First(model)
	return model, err
}
