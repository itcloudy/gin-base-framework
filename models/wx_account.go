package models

import "github.com/jinzhu/gorm"

type WxAppId struct {
	BaseModel
	OpenId string `json:"open_id" form:"open_id"` //微信id
}

func (model *WxAppId) Create(db *gorm.DB) (*WxAppId, error) {
	err := db.Create(model).Error
	db.First(model)
	return model, err
}
