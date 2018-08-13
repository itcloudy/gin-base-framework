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

type userModel struct {
	OpenId string `json:"open_id"`
	Name   string `json:"name"`
	Head   string `json:"head"`
}

type openCode struct {
	SessionKey string `json:"session_key"`
	ExpiresIn  int    `json:"expires_in"`
	OpenId     string `json:"openid"`
}

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
		code  int
	)
	jsCode := c.Query("code")
	appId := common.WeChatInfo.AppID
	secret := common.WeChatInfo.Secret
	url := common.WeChatInfo.OpenIdUrl
	requestUrl := common.StringsJoin(url, "&appid=", appId, "&secret=", secret, "&js_code=", jsCode)
	response, _ := http.Get(requestUrl)

	body, _ := ioutil.ReadAll(response.Body)
	returnMap := &openCode{}
	defer response.Body.Close()
	if response.StatusCode == 200 {

		err = json.Unmarshal(body, returnMap)
		if err != nil {
			code = common.GET_OPEN_ID_FAILED
			common.GenResponse(c, code, string(body), "wechat openid get failed")
			return
		}
		model.OpenId = returnMap.OpenId

		r, _, code := services.GetOpenId(model.OpenId)
		if r.ID == 0 {
			_, err, code = services.OpenIdCreate(&model)
		}
		if err != nil {
			common.GenResponse(c, code, string(body), err.Error())
			return
		} else {
			common.GenResponse(c, common.SUCCESSED, returnMap, "success")
			return
		}

	} else {
		common.GenResponse(c, code, string(body), "wechat openid get failed")
	}
}
