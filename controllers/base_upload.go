package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itcloudy/gin-base-framework/common"
	"github.com/itcloudy/gin-base-framework/storage"
	"strings"
	"fmt"
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
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, "only can upload "+common.UPLOAD_FILE_MIME+" image")
			return
		}
		// 限制图片大小
		size := 1024 * 1024 * common.Image.Size
		if header.Size > int64(size) {
			common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, fmt.Sprintf("size of image bigger than %d",common.Image.Size))
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
