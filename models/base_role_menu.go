package models

import "github.com/jinzhu/gorm"

type RoleMenu struct {
	BaseModel
	Role   Role `json:"role" yaml:"role"`       // 角色
	RoleID int  `json:"role_id" yaml:"role_id"` // 角色ID
	Menu   Menu `json:"menu" yaml:"menu"`       // 菜单
	MenuID int  `json:"menu_id" yaml:"menu_id"` // 菜单ID
}

func (rm *RoleMenu) Create(db *gorm.DB) (*RoleMenu, error) {

	err := db.Create(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *RoleMenu) Update(db *gorm.DB) (*RoleMenu, error) {
	err := db.Model(rm).Updates(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *RoleMenu) Delete(db *gorm.DB) (*RoleMenu, error) {
	err := db.Delete(rm).Error
	return rm, err
}
