package services

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
)

//ResourcePageCreate menu create
func ResourcePageCreate(rp *models.ResourcePage) (*models.ResourcePage, error) {
	var (
		err error
	)

	rp, err = rp.Create(common.DB)
	return rp, err
}
