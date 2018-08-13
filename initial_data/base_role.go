package initial_data

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"

	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/daemons"
)

type roleBase struct {
	Roles []*baseRoleInfo `yaml:"roles"`
}
type baseRoleInfo struct {
	Name           string             `yaml:"name"`
	MenuUniqueTags []string           `yaml:"menus"`
	Code           string             `yaml:"code"`
	ApiList        []models.SystemApi `yaml:"api_list"`
}

func InitBaseRole() {
	if len(services.GetAllRoleFromDB()) > 0 {
		return
	}
	roleIds := initBaseRole()
	if len(roleIds) > 0 {
		common.CasbinRoleIds = roleIds
	}

}
func AddRolePolicy(roleIds []int) {
	var (
		policies []common.PolicyAction
		err      error
	)
	for _, roleId := range roleIds {
		// 获得该角色所有的接口
		var (
			roleApis []*models.RoleApi
		)
		roleApis, err, _ = services.GetRoleAPIsByRoleId(roleId)
		for _, roleApi := range roleApis {
			var policy common.PolicyAction
			policy.Address = roleApi.SystemApi.Address
			policy.Method = roleApi.SystemApi.Method
			policy.PType = fmt.Sprintf("role_%d", roleId)
			policies = append(policies, policy)

		}

	}
	if len(policies) > 0 {
		tx := common.DB.Begin()
		defer func() {
			if err == nil {
				err = tx.Commit().Error
			} else {
				err = tx.Rollback().Error
			}

		}()

		if daemons.RoleSystemApiEnforcerDaemon(policies) != nil {
			log.Errorf("role casbin create  failed: %s ", err)
			log.Flush()
			os.Exit(-1)
		}
	}

	// 清空
	common.CasbinRoleIds = nil
}
func initBaseRole() []int {
	var (
		baseR   roleBase
		err     error
		roleIds []int
	)

	filePath := path.Join(common.WorkSpace, "role_data.yml")
	roleData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("role init file read failed: %s", err)
		log.Flush()
		os.Exit(-1)
	}
	err = yaml.Unmarshal(roleData, &baseR)
	if err != nil {
		log.Errorf("role init data parse failed: %s", err)
		log.Flush()
		os.Exit(-1)

	}
	tx := common.DB.Begin()
	defer func() {
		if err == nil {
			err = tx.Commit().Error
		} else {
			err = tx.Rollback().Error
		}

	}()
	for _, role := range baseR.Roles {
		// 创建角色
		var (
			mRole  models.Role
			reRole *models.Role
			err    error
		)
		mRole.Name = role.Name
		mRole.Code = role.Code
		reRole, err = mRole.Create(tx)
		if err == nil {
			roleIds = append(roleIds, reRole.ID)
			// 创建角色拥有的菜单
			for _, tag := range role.MenuUniqueTags {
				var menu *models.Menu
				var roleMenu models.RoleMenu
				menu, err, _ = services.GetMenuByUniqueTag(tag)
				if err != nil {
					log.Errorf("get menu by unique tag:%s failed: %s ", tag, err)
					log.Flush()
					os.Exit(-1)
				} else {
					roleMenu.MenuID = menu.ID
					roleMenu.RoleID = reRole.ID
					_, err = roleMenu.Create(tx)
					if err != nil {
						log.Errorf("create role menu failed : %s ,role: %+v,menu: %+v", err, reRole, menu)
						log.Flush()
						os.Exit(-1)
					}
				}

			}
			// 创建角色拥有的API
			for _, api := range role.ApiList {

				var (
					systemAPI *models.SystemApi
					roleApi   models.RoleApi
				)
				systemAPI, err, _ = services.GetSystemAPIByThreeParams(api.Name, api.Method, api.Address)
				if err != nil {
					log.Errorf("get system api failed: %+v, %s ,api: %s ,%s, %s", role, err, api.Name, api.Method, api.Address)
					log.Flush()
					os.Exit(-1)
				}
				roleApi.SystemApiID = systemAPI.ID
				roleApi.RoleID = reRole.ID
				_, err = roleApi.Create(tx)
				if err != nil {
					log.Errorf("role Api create  failed: %s ", err)
					log.Flush()
					os.Exit(-1)
				} else {

				}

			}
		}

	}

	return roleIds
}
