package controllers

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/hexiaoyun128/gin-base-framework/storage"
)

// @tags  通用
// @Description 图片上传
// @Summary 图片上传
// @Accept  multipart/form-data
// @Produce  json
// @Param Authorization header string true "Token"
// @Param file query file true "上传文件"
// @Success 200 {string} json ""
// @Router /auth/image_upload [post]
func ImageUpload(c *gin.Context) {
	var (
		filePath  string
		sufixList []string
		sufix     string
	)
	if file, header, err := c.Request.FormFile("image_file"); file != nil && header != nil {
		// 过滤mime类型
		sufixList = strings.Split(header.Filename, ".")
		sufix = sufixList[(len(sufixList) - 1)]
		if strings.Index(common.UPLOAD_FILE_MIME, sufix) == -1 {
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, "只允许上传 "+common.UPLOAD_FILE_MIME+"的图片")
			return
		}
		// 限制图片大小
		size := 1024 * 1024 * common.Image.Size
		if header.Size > int64(size) {
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, "图片太大！")
			return
		}
		if filePath, err = storage.FileLocalStorage(file, header.Filename); err != nil {
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, err, err.Error())
			return
		}
	} else {
		common.GenResponse(c, common.UPLOAD_FILE_RESROUCE_ERR, err, err.Error())
		return
	}
	common.GenResponse(c, common.SUCCESSED, filePath, "success")
}


func MultiUpload(c *gin.Context) {
	/*form, _ := c.MultipartForm()
	files := form.File["upload[]"]*/
}
