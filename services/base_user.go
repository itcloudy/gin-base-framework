package services

import (
	"errors"
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/daemons"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"gopkg.in/go-playground/validator.v9"
)

func GetUserById(id int) (*models.User, error) {

	var (
		user models.User
		err  error
	)

	err = common.DB.First(&user, "id = ?", id).Related(&user.Roles, "Roles").Error
	return &user, err
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := common.DB.First(&user, "name = ?", name).Related(&user.Roles, "Roles").Error
	return &user, err
}

func GetUserByNameOrEmail(name string) (*models.User, error) {
	var user models.User
	err := common.DB.First(&user, "name = ? OR email = ?", name, name).Related(&user.Roles, "Roles").Error
	return &user, err
}

func GetUserByOpenId(user models.User) (*models.User, error) {
	var (
		err error
	)
	db := common.DB
	err = db.First(&user, "open_id = ?", user.OpenId).Related(&user.Roles, "Roles").Error
	if err != nil {
		return user.Create(common.DB)
	}
	return &user, err
}
func UserCreate(user *models.User) (*models.User, error) {
	var (
		validate *validator.Validate
		err      error
	)
	validate = validator.New()
	err = validate.Struct(user)
	if err != nil {
		return nil, err
	}

	user.Pwd = common.Md5(user.Password)

	user, err = user.Create(common.DB)


	return user, err
}
func CheckUser(name, password string) (*models.User, error) {
	user, _ := GetUserByNameOrEmail(name)
	if user.Pwd == common.Md5(password) {
		return user, nil
	} else {
		return nil, errors.New("name or password invalid")
	}
}

//UserUpdate role create
func UserUpdate(user *models.User, admin bool) (*models.User, error) {
	var (
		err                error
		groupPolicyActions []common.GroupPolicyAction
	)
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
	user, err = user.Update(tx)




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

				role, err = GetRoleById(roleId)
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
				tx.Model(user).Association("Roles").Append(roleSlice)
			}
		}
	}
	return user, err
}
func GetAllUser(whereQuery string, whereArgs []interface{}, order string, page, limit int) ([]*models.User, error) {
	var (
		err      error
		userList []*models.User
	)
	db := common.DB
	if whereQuery != "" {
		db = db.Where(whereQuery, whereArgs...)
	}
	err = db.Find(&userList).Offset((page - 1) * limit).Limit(limit).Order(order).Error
	return userList, err
}
