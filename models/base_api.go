package models

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/jinzhu/gorm"
	"strings"
)

type SystemApi struct {
	BaseModel
	Name    string `json:"name" yaml:"name" gorm:"index"`       // API名称
	Address string `json:"address" yaml:"address" gorm:"index"` // API地址
	Method  string `json:"method" yaml:"method" gorm:"index"`   // API请求方法
	Display string `json:"display" yaml:"-" gorm:"unique"`      // 显示
}

func (rm *SystemApi) Create(db *gorm.DB) (*SystemApi, error) {
	rm.Method = strings.ToUpper(rm.Method)
	rm.Display = common.StringsJoin("[", rm.Method, "] ", rm.Name, " (", rm.Address, ")")
	err := db.Create(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *SystemApi) Update(db *gorm.DB) (*SystemApi, error) {
	rm.Method = strings.ToUpper(rm.Method)
	rm.Display = common.StringsJoin("[", rm.Method, "] ", rm.Name, " (", rm.Address, ")")
	err := db.Model(rm).Updates(rm).Error
	db.First(rm)
	return rm, err
}

func (rm *SystemApi) Delete(db *gorm.DB) (*SystemApi, error) {
	err := db.Delete(rm).Error
	return rm, err
}
