package models

import (
	"errors"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/jinzhu/gorm"
)

type User struct {
	BaseModel
	Name            string  `form:"name" json:"name" `                                                             // 用户名
	Alias           string  `json:"alias" form:"alias"`                                                            // 用户昵称
	Head            string  `json:"head" form:"head"`                                                              // 用户头像
	Email           string  `json:"email" validate:"email" form:"email"`                                           // 用户邮箱
	Mobile          string  `json:"mobile" `                                                                       // 用户手机号
	Password        string  `gorm:"-" json:"password" form:"password" validate:"password,eqfield=ConfirmPassword"` // 用户密码,创建时的明文
	ConfirmPassword string  `gorm:"-" json:"confirm_password" form:"confirm_password"`                             // 用户确认密码,创建时的明文
	Pwd             string  `json:"-"`                                                                             // 用户明码，数据保存
	Roles           []*Role `json:"roles" gorm:"many2many:role_user;"`                                             // 拥有角色
	RoleList        []int   `gorm:"-" json:"role_list,omitempty"`                                                  // 拥有的角色ID列表
	OpenId          string  `json:"openid" form:"openid" gorm:"unique"`                                            // 微信openID
	IsActive        bool    `gorm:"default:true" json:"active" form:"active"`                                      // 是否有效
	IsAdmin         bool    `gorm:"default:false" json:"is_admin" form:"is_admin"`                                 // 是否为管理员
}

func (user *User) Create(db *gorm.DB) (*User, error) {
	user.IsAdmin = false
	var (
		model User
		err   error
	)
	if user.Mobile != "" {
		err = db.First(&model, "mobile = ?", user.Mobile).Error
		if err == nil {
			err := errors.New("mobile has registered!")
			return nil, err
		}
	}

	err = db.Create(user).Error
	db.First(user)
	return user, err
}

func (user *User) Update(db *gorm.DB) (*User, error) {
	user.IsAdmin = false
	var (
		model User
		err   error
	)
	if user.Mobile != "" {
		err = db.First(&model, "mobile = ?", user.Mobile).Error
		if err == nil {
			err := errors.New("mobile has registered!")
			return nil, err
		}
	}
	err = db.Model(user).Updates(user).Error
	db.First(user)
	return user, err
}
func (user *User) ActiveAction(db *gorm.DB) (*User, error) {
	err := db.Model(user).Updates(map[string]interface{}{
		"active": user.IsActive,
	}).Error
	return user, err
}
func (user *User) SetPwd(db *gorm.DB) (*User, error) {
	pwd := common.SHA256(user.Password)
	err := db.Model(user).Updates(map[string]interface{}{
		"pwd": pwd,
	}).Error
	return user, err
}

func (user *User) AdminAction(db *gorm.DB) (*User, error) {
	err := db.Model(user).Updates(map[string]interface{}{
		"is_admin": user.IsAdmin,
	}).Error
	return user, err
}
