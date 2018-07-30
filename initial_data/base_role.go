package initial_data

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type roleBase struct {
	Roles []*baseRoleInfo `yaml:"roles"`
}
type baseRoleInfo struct {
	Name     string   `yaml:"name"`
	MenuList []string `yaml:"menu_list"`
	ApiList  []string `yaml:"api_list"`
}

func InitBaseRole() {
	var baseR roleBase
	wr, _ := os.Getwd()
	filePath := path.Join(wr, "conf", "base_role.yml")
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
}
