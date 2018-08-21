package initial_data

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"

	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"go.uber.org/zap"
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
		common.Logger.Error("api init file read failed", zap.Error(err))
		fmt.Printf("api init file read failed: %s", err)

		os.Exit(-1)
	}
	err = yaml.Unmarshal(apiData, &systemapis)
	if err != nil {
		common.Logger.Error("system api init data parse failed", zap.Error(err))
		fmt.Printf("system api init data parse failed: %s", err)

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
			common.Logger.Error("system api create failed ", zap.Error(err))

			os.Exit(-1)
		}
	}
}
