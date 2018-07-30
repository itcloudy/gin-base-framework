package controllers

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/gin-gonic/gin"
	"github.com/hexiaoyun128/gin-base-framework/storage"
)

// @tags  文件上传
// @Description 文件上传
// @Summary 文件上传
// @Accept  multipart/form-data
// @Produce  json
// @Param Authorization header string true "Token"
// @Param file query file true "上传文件"
// @Success 200 {string} json ""
// @Router /auth/file_upload [post]
func FileUpload(c *gin.Context) {
	var (
		filePath string
	)
	if file, header, err := c.Request.FormFile("file"); file != nil && header != nil {
		if filePath, err = storage.FileLocalStorage(file, header.Filename); err != nil {
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, err, err.Error())
			return
		}
	} else {
		common.GenResponse(c, common.UPLOAD_FILE_RESROUCE_ERR, err, err.Error())
		return
	}
	filePath = common.StringsJoin(common.UPLOAD_FILE_URL, filePath)
	common.GenResponse(c, common.SUCCESSED, filePath, "success")
}

func MultiUpload(c *gin.Context) {
	/*form, _ := c.MultipartForm()
	files := form.File["upload[]"]*/
}
