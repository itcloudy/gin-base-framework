package models

import (
	"github.com/jinzhu/gorm"
)

type Role struct {
	BaseModel
	Name       string      `form:"name" gorm:"not null;unique_index" json:"name" binding:"required"` // 角色名称
	Code       string      `json:"code"`                                                             // 角色编码
	InheritIds []int       `gorm:"-"`                                                                // 所继承的角色ID
	Inherits   []*Role     `yaml:"inherits" json:"inherits"`                                         // 继承的角色
	RoleMenus  []*RoleMenu `json:"role_menus" yaml:"role_menus"`                                     // 角色拥有的菜单
}

func (role *Role) Create(db *gorm.DB) (*Role, error) {
	err := db.Create(role).Error
	db.First(role)
	return role, err
}

func (role *Role) Update(db *gorm.DB) (*Role, error) {
	err := db.Model(role).Updates(role).Error
	db.First(role)
	return role, err
}

func (role *Role) Delete(db *gorm.DB) (*Role, error) {
	err := db.Delete(role).Error
	return role, err
}
