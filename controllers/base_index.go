package controllers

import (
	"fmt"
	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/middles"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"github.com/gin-gonic/gin"
)

func IndexGet(c *gin.Context) {

	common.GenResponse(c, common.SUCCESSED, "ok", "success")
}

// @tags  Token刷新
// @Description Token刷新
// @Summary Token刷新
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Success 200 {string} json "{"code":200,"data":"Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjc1MzIxNjI5OTEsIk5hbWUiOiJhZG1pbiIsIlJvbGUiOm51bGwsIlVzZXJJZCI6MSwiSXNBZG1pbiI6dHJ1ZX0.je9O18vV-E-CENlXnGUk1eztLChZrd9kqsfyCLjg02PE_VNFu6UUlg952qveuJ8yL1yPe-GEdA-FYqMFrqWk1ekQAbDphbO15oD0GpvHZhiTGL3XunNFU_LudLvoOuNEAnEACHxklfMY1J37jesrehhoqxA5pcp0ushINUprsds","message":"success"}"
// @Router /auth/refresh [get]
func RefreshToken(c *gin.Context) {
	var (
		user *models.User
		err  error
		code int
	)
	user, err, code = services.GetUserById(c.GetInt(common.LOGIN_USER_ID))

	if err == nil {
		var roleList []string
		for _, role := range user.Roles {
			roleList = append(roleList, fmt.Sprintf("role_%d", role.ID))
		}
		token := middles.GenerateJWT(user.Name, roleList, user.ID, user.IsAdmin)
		common.GenResponse(c, common.SUCCESSED, token, "success")
	} else {
		common.GenResponse(c, code, nil, " refresh token failed")
	}
}
