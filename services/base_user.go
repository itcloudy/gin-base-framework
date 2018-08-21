package services

import (
	"errors"
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/daemons"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/utils"
	"gopkg.in/go-playground/validator.v9"
)

func GetUserById(id int) (*models.User, error, int) {

	var (
		user models.User
		err  error
		code int
	)
	code = common.SUCCESSED

	if err = common.DB.First(&user, "id = ?", id).Related(&user.Roles, "Roles").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &user, err, code
}

func GetUserByName(name string) (*models.User, error, int) {
	var (
		user models.User
		err  error
		code int
	)
	if err = common.DB.First(&user, "name = ?", name).Related(&user.Roles, "Roles").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &user, err, code
}

func GetUserByMobile(mobile string) (*models.User, error, int) {
	var (
		user models.User
		err  error
		code int
	)
	if err = common.DB.First(&user, "mobile = ?", mobile).Related(&user.Roles, "Roles").Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}

	return &user, err, code
}

func GetUserByOpenId(user *models.User) (*models.User, error, int) {
	var (
		err  error
		code int
	)
	db := common.DB
	err = db.First(&user, "open_id = ?", user.OpenId).Error
	if err != nil {
		//return user.Create(common.DB)
		//创建用户的同时添加权限
		return UserCreate(user)
	}
	db.Model(&user).Related(&user.Roles, "Roles")
	return user, err, code
}
func UserCreate(user *models.User) (*models.User, error, int) {
	var (
		validate *validator.Validate
		err      error
		code     int
	)
	code = common.SUCCESSED

	if user.OpenId == "" {
		validate = validator.New()
		validate.RegisterValidation("password", utils.ValidatePassword)
		err = validate.Struct(user)
		if err != nil {
			code = common.DATA_VALIDATE_ERR
			return nil, err, code
		}
	}

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}

	}()

	//  判断手机号码是否重复
	if user.Mobile != "" {
		if u, err, _ := GetUserByMobile(user.Mobile); err == nil {
			if u.ID > 0 {
				code = common.DATA_VALIDATE_ERR
				return nil, errors.New("mobile repeat"), code
			}
		}
	}

	user, err = user.Create(tx)
	if user.Password != "" {
		user.SetPwd(tx)
	}
	//添加通用用户权限 所有用户都有ordinary角色
	var roleSlice []models.Role
	if ordinary, e, co := GetRoleByCode("ordinary"); e == nil {
		roleSlice = append(roleSlice, *ordinary)
	} else {
		return nil, e, co
	}

	tx.Model(user).Association("Roles").Append(roleSlice)

	return user, err, code
}

func CheckUser(mobile, password string) (user *models.User, err error, code int) {
	if user, err, code = GetUserByMobile(mobile); err != nil {
		return
	}
	if user.Pwd == common.SHA256(password) {
		db := common.DB
		db.Model(user).Related(&user.Roles, "Roles")
		return
	} else {
		return nil, errors.New("name or password invalid"), common.NAME_PASSWORD_INVALID
	}
}

//UserUpdate role create
func UserUpdate(user *models.User, admin bool) (reModel *models.User, err error, code int) {
	var (
		groupPolicyActions []common.GroupPolicyAction
	)
	code = common.SUCCESSED

	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}
		if err == nil {
			go daemons.UserOrRoleEnforcerDaemon(groupPolicyActions)
		}

	}()
	//判断手机号码是否重复
	if user.Mobile != "" {
		if reModel, err, code = GetUserByMobile(user.Mobile); err == nil {
			if reModel.ID > 0 && reModel.ID != user.ID {
				err = errors.New("mobile repeat")
				code = common.DATA_VALIDATE_ERR
				return
			}
		}
	}
	if user, err = user.Update(tx); err != nil {
		code = common.DB_UPDATE_ERR
		return
	}

	if err == nil {
		if admin {
			var roleSlice []models.Role
			tx.Model(user).Related(&user.Roles, "Roles")
			if len(user.Roles) > 0 {
				for _, ro := range user.Roles {
					var gpa common.GroupPolicyAction
					gpa.Action = "delete"
					gpa.Role = fmt.Sprintf("role_%d", ro.ID)
					gpa.UserOrRole = fmt.Sprintf("user_%d", user.ID)
					groupPolicyActions = append(groupPolicyActions, gpa)

				}
				tx.Model(user).Association("Roles").Delete(user.Roles)
			}

			for _, roleId := range user.RoleList {
				var (
					role *models.Role
				)

				role, err, _ = GetRoleById(roleId)
				if err != nil {
					break
				} else {
					roleSlice = append(roleSlice, *role)
				}

			}
			if err == nil {
				for _, ro := range roleSlice {
					var gpa common.GroupPolicyAction
					gpa.Action = "add"
					gpa.Role = fmt.Sprintf("role_%d", ro.ID)
					gpa.UserOrRole = fmt.Sprintf("user_%d", user.ID)
					groupPolicyActions = append(groupPolicyActions, gpa)

				}
				// 添加通用用户权限
				if ordinary, e, _ := GetRoleByCode("ordinary"); e == nil {
					roleSlice = append(roleSlice, *ordinary)
				}
				tx.Model(user).Association("Roles").Append(roleSlice)
			}
		}
	}
	return
}
func GetAllUser(whereQuery string, whereArgs []interface{}, order string, page, limit int) (userList []*models.User, total int, err error, code int) {

	db := common.DB
	if whereQuery != "" {
		db = db.Where(whereQuery, whereArgs...)
	}
	code = common.SUCCESSED
	if err = db.Where(whereQuery, whereArgs...).Offset((page - 1) * limit).Limit(limit).Order(order).Find(&userList).Error; err == nil {
		db.Model(models.User{}).Where(whereQuery, whereArgs...).Count(&total)
	} else {
		code = common.DB_RECORD_NOT_FOUND

	}
	return
}

//通过openid获取用户
func GetUser(openId string) (*models.User, error, int) {
	var (
		model models.User
		err   error
		code  int
	)
	db := common.DB
	err = db.First(&model, "open_id = ?", openId).Error
	if err != nil {
		code = common.DB_RECORD_NOT_FOUND
		return nil, err, code
	}
	code = common.SUCCESSED
	return &model, err, code
}

//根据openid更新用户名与头像
func UpdateUser(openId, name, head string) (*models.User, error, int) {
	var (
		model models.User
		err   error
		code  int
	)
	db := common.DB
	err = db.Model(&model).Where("open_id = ?", openId).Updates(models.User{Name: name, Head: head}).Error
	if err != nil {
		code = common.DB_UPDATE_ERR
		return &model, err, code
	}
	code = common.SUCCESSED
	return &model, err, code
}
