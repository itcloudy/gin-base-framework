package services

import (
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"strconv"
	"strings"
)

//MenuCreate menu create
func MenuCreate(menu *models.Menu) (*models.Menu, error) {
	var (
		err    error
		parent *models.Menu
	)
	if menu.ParentID > 0 {
		if parent, err = GetMenuById(menu.ParentID); err == nil {
			if parent.Tree != "" {
				menu.Tree = fmt.Sprintf("%s-%d", parent.Tree, parent.ID)
			} else {
				menu.Tree = fmt.Sprintf("%d", parent.ID)
			}
		}
	}
	menu, err = menu.Create(common.DB)
	return menu, err

}

//MenuUpdate menu create
func MenuUpdate(menu *models.Menu) (*models.Menu, error) {
	var (
		err    error
		parent *models.Menu
	)
	if menu.ParentID > 0 {
		if parent, err = GetMenuById(menu.ParentID); err == nil {
			if parent.Tree != "" {
				menu.Tree = fmt.Sprintf("%s-%d", parent.Tree, parent.ID)
			} else {
				menu.Tree = fmt.Sprintf("%d", parent.ID)
			}
		}
	}
	menu, err = menu.Update(common.DB)

	return menu, err

}

//GetMenuById get menu by id
func GetMenuById(id int) (*models.Menu, error) {
	var (
		menu models.Menu
		err  error
	)
	err = common.DB.First(&menu, "id = ?", id).Error
	return &menu, err
}

func GetMenuByName(name string) (*models.Menu, error) {
	var menu models.Menu
	err := common.DB.First(&menu, "name = ?", name).Error
	return &menu, err
}

func GetMenuByRoute(route string) (*models.Menu, error) {
	var menu models.Menu
	err := common.DB.First(&menu, "route = ?", route).Error
	return &menu, err
}

func GetMenuByUserRoleIds(roleIds []int, isAdmin bool) ([]*models.Menu, error) {
	var (
		menus       []*models.Menu
		roleMenus   []*models.RoleMenu
		err         error
		menuIds     []int
		allMenuIds  []int
		returnMenus []*models.Menu
	)
	menuMaps := make(map[int]*models.Menu)
	db := common.DB
	if isAdmin {
		err = db.Find(&menus).Error
	} else {
		err = db.Where("role_id IN (?)", roleIds).Find(&roleMenus).Error
		if len(roleMenus) > 0 {
			for _, rm := range roleMenus {
				menuIds = append(menuIds, rm.MenuID)
			}
		}
		err = db.Where("id IN (?)", menuIds).Find(&menus).Error
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
		err = db.Where("id IN (?)", menuIds).Find(&menus).Error
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

	return returnMenus, err
}

//GetMenuByUserID get user menu tree by user id
func GetMenuByUserID(userId int) ([]*models.Menu, error) {
	var (
		roleIds []int
		user    models.User
	)
	db := common.DB
	user.ID = userId
	db.First(&user)
	for _, role := range user.Roles {
		roleIds = append(roleIds, role.ID)
	}
	return GetMenuByUserRoleIds(roleIds, user.IsAdmin)

}
