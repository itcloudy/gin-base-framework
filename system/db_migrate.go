package system

import (
	"github.com/itcloudy/gin-base-framework/models"
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB) {
	//migrate base model to database table
	db.AutoMigrate(&models.User{}, &models.Menu{}, &models.SystemApi{}, &models.RoleApi{},
		&models.Role{}, &models.RoleMenu{})
	//application model to database table
	db.AutoMigrate(
		&models.WxAppId{},
	)

}
