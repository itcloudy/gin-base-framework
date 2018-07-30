package models

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Menu struct {
	BaseModel
	Parent    *Menu   `json:"parent,omitempty"`                                   // 上级菜单
	ParentID  int     `json:"parent_id"`                                          // 上级菜单ID
	Name      string  `json:"name" gorm:"index" yaml:"name" validate:"required" ` // 菜单名称
	Route     string  `json:"route,omitempty" yaml:"route"`                       // 菜单路由
	Component string  `json:"component,omitempty" yaml:"component"`               // 菜单组件
	Icon      string  `json:"icon,omitempty" yaml:"icon" validate:"required"`     // 菜单样式类
	Sequence  int     `json:"sequence" yaml:"sequence" validate:"required"`       // 菜单顺序
	Tree      string  `json:"-" yaml:"tree" `                                     // 菜单继承树
	Children  []*Menu `json:"children,omitempty" yaml:"children"`                 // 子菜单
	UniqueTag string  `json:"unique_tag" gorm:"unique_index" validate:"required"` // 菜单唯一标识

}

func (menu *Menu) Create(db *gorm.DB) (*Menu, error) {
	if menu.UniqueTag == "" {
		m, _ := uuid.NewV4()
		menu.UniqueTag = m.String()
	}
	err := db.Create(menu).Error
	db.First(menu)
	return menu, err
}

func (menu *Menu) Update(db *gorm.DB) (*Menu, error) {
	err := db.Model(menu).Updates(menu).Error
	db.First(menu)
	return menu, err
}

func (menu *Menu) Delete(db *gorm.DB) (*Menu, error) {
	err := db.Delete(menu).Error
	return menu, err
}

//CreateOrUpdate only create or update one record,ignore children information
func (menu *Menu) CreateOrUpdate(db *gorm.DB) (*Menu, error) {
	var mem = Menu{}
	db.Where("name = ?", menu.Name).First(&mem)
	if mem.ID > 0 {
		mem.Parent = menu.Parent
		mem.ParentID = menu.ParentID
		mem.Name = menu.Name
		return mem.Update(db)

	} else {

		return menu.Create(db)
	}
}
