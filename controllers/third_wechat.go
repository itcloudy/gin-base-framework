package controllers

import (
	"encoding/json"

	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// @tags 微信
// @Description 用户用户openid
// @Summary 用户用户openid
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param code query string true "微信Code"
// @Success 200 {string} json  "{}"
// @Router /auth/banner_group [post]
func GetOpenId(c *gin.Context) {

	var (
		model models.WxAppId
		err   error
	)
	jsCode := c.Query("code")
	appId := common.WeChatInfo.AppID
	secret := common.WeChatInfo.Secret
	url := common.WeChatInfo.OpenIdUrl
	requestUrl := common.StringsJoin(url, "&appid=", appId, "&secret=", secret, "&js_code=", jsCode)
	response, _ := http.Get(requestUrl)

	body, _ := ioutil.ReadAll(response.Body)
	var returnMap map[string]string
	defer response.Body.Close()
	if response.StatusCode == 200 {
		json.Unmarshal(body, &returnMap)
		model.OpenId = returnMap["openid"]
		if model.OpenId == "" {
			common.GenResponse(c, common.FAILED, body, "wechat openid get failed")
			return
		}
		r, _ := services.GetOpenId(model.OpenId)
		if r.ID == 0 {
			_, err = services.OpenIdCreate(&model)
		}
		if err != nil {
			common.GenResponse(c, common.FAILED, body, err.Error())
			return
		} else {
			common.GenResponse(c, common.SUCCESSED, returnMap, "success")
			return
		}

	} else {
		common.GenResponse(c, common.FAILED, body, "wechat openid get failed")
	}
}
