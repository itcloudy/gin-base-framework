package initial_data

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"

	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
)

type systemApis struct {
	Apis []models.SystemApi `yaml:"apis" json:"apis"`
}

//InitApi
func InitApi() {

	var systemapis *systemApis

	filePath := path.Join(common.WorkSpace, "api_data.yml")
	apiData, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("api init file read failed: %s", err)
		fmt.Printf("api init file read failed: %s", err)
		log.Flush()
		os.Exit(-1)
	}
	err = yaml.Unmarshal(apiData, &systemapis)
	if err != nil {
		log.Errorf("system api init data parse failed: %s", err)
		fmt.Printf("system api init data parse failed: %s", err)
		log.Flush()
		os.Exit(-1)
	}
	if len(services.GetAllSystemApiFromDB()) == 0 {
		insertApis(systemapis)
	}

}
func insertApis(apis *systemApis) {
	var (
		err error
	)
	for _, a := range apis.Apis {
		_, err, _ = services.SystemApiCreate(&a)
		if err != nil {
			log.Errorf("system api create failed : %s,%+v", err, a)
			log.Flush()
			os.Exit(-1)
		}
	}
}
