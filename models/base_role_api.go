package models

import "github.com/jinzhu/gorm"

type RoleApi struct {
	BaseModel
	Role        Role       `json:"role" yaml:"role"`                   // 角色
	RoleID      int        `json:"role_id" yaml:"role_id"`             // 角色ID
	SystemApiID int        `json:"system_api_id" yaml:"system_api_id"` // 接口ID
	SystemApi   *SystemApi `json:"system_api" yaml:"system_api"`       // 接口
}

func (rm *RoleApi) Create(db *gorm.DB) (*RoleApi, error) {
	err := db.Create(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *RoleApi) Update(db *gorm.DB) (*RoleApi, error) {
	err := db.Model(rm).Updates(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *RoleApi) Delete(db *gorm.DB) (*RoleApi, error) {
	err := db.Delete(rm).Error
	return rm, err
}
