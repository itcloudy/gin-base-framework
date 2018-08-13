package services

import (
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"strconv"
	"strings"
)

//MenuCreate menu create
func MenuCreate(menu *models.Menu) (*models.Menu, error, int) {
	var (
		err    error
		parent *models.Menu
		code   int
	)
	code = common.SUCCESSED
	if menu.ParentID > 0 {
		if parent, err, _ = GetMenuById(menu.ParentID); err == nil {
			if parent.Tree != "" {
				menu.Tree = fmt.Sprintf("%s-%d", parent.Tree, parent.ID)
			} else {
				menu.Tree = fmt.Sprintf("%d", parent.ID)
			}
		}
	}
	if menu, err = menu.Create(common.DB); err != nil {
		code = common.DB_INSERT_ERR
	}
	return menu, err, code

}

//MenuUpdate menu create
func MenuUpdate(menu *models.Menu) (*models.Menu, error, int) {
	var (
		err    error
		parent *models.Menu
		code   int
	)
	code = common.SUCCESSED
	if menu.ParentID > 0 {
		if parent, err, _ = GetMenuById(menu.ParentID); err == nil {
			if parent.Tree != "" {
				menu.Tree = fmt.Sprintf("%s-%d", parent.Tree, parent.ID)
			} else {
				menu.Tree = fmt.Sprintf("%d", parent.ID)
			}
		}
	}
	if menu, err = menu.Update(common.DB); err != nil {
		code = common.DB_UPDATE_ERR
	}

	return menu, err, code

}

//GetMenuById get menu by id
func GetMenuById(id int) (*models.Menu, error, int) {
	var (
		menu models.Menu
		err  error
		code int
	)
	code = common.SUCCESSED
	if err = common.DB.First(&menu, "id = ?", id).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &menu, err, code
}

func GetMenuByName(name string) (*models.Menu, error, int) {
	var (
		menu models.Menu
		err  error
		code int
	)
	code = common.SUCCESSED
	if err = common.DB.First(&menu, "name = ?", name).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}
	return &menu, err, code
}

func GetMenuByUserRoleIds(roleIds []int, isAdmin bool) (returnMenus []*models.Menu, err error, code int) {
	var (
		menus      []*models.Menu
		roleMenus  []*models.RoleMenu
		menuIds    []int
		allMenuIds []int
	)
	menuMaps := make(map[int]*models.Menu)
	code = common.SUCCESSED

	db := common.DB
	if isAdmin {
		if err = db.Find(&menus).Error; err != nil {
			code = common.SYSTEM_HAS_NO_MENUS
			return
		}
	} else {
		if err = db.Where("role_id IN (?)", roleIds).Find(&roleMenus).Error; err != nil {
			code = common.ROLE_HAS_NO_MENUS
			return
		}
		if len(roleMenus) > 0 {
			for _, rm := range roleMenus {
				menuIds = append(menuIds, rm.MenuID)
			}
		}
		if err = db.Where("id IN (?)", menuIds).Find(&menus).Error; err != nil {
			code = common.DB_RECORD_NOT_FOUND
			return
		}
		for _, m := range menus {
			tree := m.Tree
			if tree != "" {
				idsStr := strings.Split(tree, "-")
				for _, idStr := range idsStr {
					if idInt, e := strconv.Atoi(idStr); e == nil {
						menuIds = append(menuIds, idInt)

					}
				}
			}
		}
		if err = db.Where("id IN (?)", menuIds).Find(&menus).Error; err != nil {
			code = common.ROLE_MENUS_TREE_ERR
			return
		}
	}
	for _, m := range menus {
		menuMaps[m.ID] = m
		allMenuIds = append(allMenuIds, m.ID)
	}
	//get menu tree
	for _, mId := range allMenuIds {
		men := menuMaps[mId]
		ParentID := men.ParentID
		if ParentID == 0 {
			returnMenus = append(returnMenus, men)
			continue
		}
		parentMenu := menuMaps[ParentID]
		if len(parentMenu.Children) == 0 {
			var children []*models.Menu
			children = append(children, men)
			parentMenu.Children = children
		} else {
			parentMenu.Children = append(parentMenu.Children, men)
		}
	}

	return
}

//GetMenuByUserID get user menu tree by user id
func GetMenuByUserID(userId int) (menus []*models.Menu, err error, code int) {
	var (
		roleIds []int
		user    models.User
	)

	db := common.DB
	user.ID = userId
	if err = db.First(&user).Error; err != nil {
		code = common.SYSTEM_HAS_NO_MENUS
		return

	}
	for _, role := range user.Roles {
		roleIds = append(roleIds, role.ID)
	}
	menus, err, code = GetMenuByUserRoleIds(roleIds, false)
	return

}

//DeleteMenupById delete menu by id
func DeleteMenupById(id int) (*models.Menu, error, int) {

	var (
		model models.Menu
		err   error
		code  int
	)
	code = common.SUCCESSED
	model.ID = id
	if err = common.DB.Delete(&model).Error; err != nil {
		code = common.DB_DELETE_ERR
	}
	return nil, err, code
}

//GetMenuByUserID get user menu tree by user id
func GetMenuByUniqueTag(uniqueTag string) (*models.Menu, error, int) {
	var (
		menu models.Menu
		err  error
		code int
	)
	code = common.SUCCESSED
	db := common.DB
	if err = db.Where("unique_tag = ?", uniqueTag).First(&menu).Error; err != nil {
		code = common.DB_RECORD_NOT_FOUND
	}

	return &menu, err, code

}
