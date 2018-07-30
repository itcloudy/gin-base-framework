package initial_data

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"os"
	"path"
)

type apiInfo struct {
	Name    string `json:"name" yaml:"name"`       // api name
	Address string `json:"address" yaml:"address"` // api address
	Method  string `json:"method" yaml:"method"`   // api method
}
type menuParam struct {
	Name      string       `yaml:"name" json:"name"`           // menu name
	Route     string       `yaml:"route" json:"route"`         // menu route
	Component string       `yaml:"component" json:"component"` // web component
	Icon      string       `yaml:"icon" json:"icon"`           // web icon
	ParentID  int          `json:"parent_id"`                  // menu parent
	Children  []*menuParam `yaml:"children" json:"children"`   // child menus
	ApiList   []*apiInfo   `yaml:"api_list" json:"api_list"`   // api list
	Sequence  int          `yaml:"sequence"`                   // display order
	UniqueTag string       `yaml:"unique_tag"`                 // 菜单唯一标识

}
type systemMenus struct {
	Menus []menuParam `yaml:"menus" json:"menus"`
}

//InitMenu
func InitMenu() {
	var ids []int
	n, _ := services.GetMenuByUserRoleIds(ids, true)
	if len(n) > 0 {
		return
	}
	var systemMs *systemMenus
	wr, _ := os.Getwd()
	filePath := path.Join(wr, "conf", "menu_data.yml")
	menuData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("menu init file read failed: %s", err)

		fmt.Printf("menu init file read failed: %s", err)

		log.Flush()
		os.Exit(-1)
	}
	err = yaml.Unmarshal(menuData, &systemMs)
	if err != nil {
		log.Errorf("menu init data parse failed: %s", err)
		fmt.Printf("menu init data parse failed: %s", err)
		log.Flush()
		os.Exit(-1)
	}
	for _, m := range systemMs.Menus {

		insertMenus(&m)
	}
}
func insertMenus(men *menuParam) {
	var (
		sysMenu models.Menu
		m       *models.Menu
		err     error
	)

	sysMenu.Name = men.Name
	sysMenu.Component = men.Component
	sysMenu.Icon = men.Icon
	sysMenu.Route = men.Route
	sysMenu.ParentID = men.ParentID
	sysMenu.Sequence = men.Sequence
	sysMenu.UniqueTag = men.UniqueTag
	m, err = services.MenuCreate(&sysMenu)
	if err != nil {
		log.Errorf("menu create failed : %s,%v", err, m)
		log.Flush()
		os.Exit(-1)
	}

	// api info
	for _, api := range men.ApiList {
		var apiInfo models.ResourceApi
		apiInfo.Name = api.Name
		apiInfo.Address = api.Address
		apiInfo.Method = api.Method
		apiInfo.MenuID = m.ID

		if _, err := services.ResourceApiCreate(&apiInfo); err != nil {
			log.Errorf("resource api create failed : %s", err)
			log.Flush()
			os.Exit(-1)
		}
	}
	for _, me := range men.Children {
		me.ParentID = m.ID
		insertMenus(me)
	}

}
