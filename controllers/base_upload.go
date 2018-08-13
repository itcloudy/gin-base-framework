package controllers

import (
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Upload(c *gin.Context) {
	var err error
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		common.GenResponse(c, common.UPLOAD_FILE_NO_EXSIT, nil, err.Error())
		return
	}
	filename := header.Filename
	out, err := os.Create("upload_files/" + filename)
	if err != nil {
		common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, err.Error())
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		common.GenResponse(c, common.UPLOAD_FILE_CREATE_ERR, nil, err.Error())

		return
	}
	common.GenResponse(c, common.SUCCESSED, common.UPLOAD_FILE_URL+filename, err.Error())

}

func MultiUpload(c *gin.Context) {
	/*form, _ := c.MultipartForm()
	files := form.File["upload[]"]*/
}
