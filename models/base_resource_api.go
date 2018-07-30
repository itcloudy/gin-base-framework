package models

import "github.com/jinzhu/gorm"

type ResourceApi struct {
	BaseModel
	Name    string `json:"name" yaml:"name"`       // API名称
	Address string `json:"address" yaml:"address"` // API地址
	Method  string `json:"method" yaml:"method"`   // API请求方法
	Menu    Menu   `json:"menu" yaml:"menu"`       // API所属菜单
	MenuID  int    `json:"menu_id" yaml:"menu_id"` // API所属菜单ID

}

func (ra *ResourceApi) Create(db *gorm.DB) (*ResourceApi, error) {
	err := db.Create(ra).Error
	db.First(ra)
	return ra, err
}

func (ra *ResourceApi) Update(db *gorm.DB) (*ResourceApi, error) {
	err := db.Model(ra).Updates(ra).Error
	db.First(ra)
	return ra, err
}

func (ra *ResourceApi) Delete(db *gorm.DB) (*ResourceApi, error) {
	err := db.Delete(ra).Error
	return ra, err
}
