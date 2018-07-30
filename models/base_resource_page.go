package models

import "github.com/jinzhu/gorm"

type ResourcePage struct {
	BaseModel
	Name   string `json:"name" yaml:"name"`       // 页面资源名称
	Sign   string `json:"sign" yaml:"sign"`       // 页面资源标记
	Type   string `json:"type" yaml:"type"`       // 标记类型 :id,clazz
	Menu   Menu   `json:"menu" yaml:"menu"`       // 所属菜单
	MenuID int    `json:"menu_id" yaml:"menu_id"` // 所属菜单ID

}

func (rp *ResourcePage) Create(db *gorm.DB) (*ResourcePage, error) {
	err := db.Create(rp).Error
	db.First(rp)
	return rp, err
}

func (rp *ResourcePage) Update(db *gorm.DB) (*ResourcePage, error) {
	err := db.Model(rp).Updates(rp).Error
	db.First(rp)
	return rp, err
}

func (ra *ResourcePage) Delete(db *gorm.DB) (*ResourcePage, error) {
	err := db.Delete(ra).Error
	return ra, err
}
