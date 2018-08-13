package controllers

import (
	"encoding/json"

	"github.com/hexiaoyun128/gin-base-framework/common"
	"github.com/hexiaoyun128/gin-base-framework/models"
	"github.com/hexiaoyun128/gin-base-framework/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"strings"
	"time"
	"strconv"
	"sort"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
	"encoding/xml"
	"bytes"
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

// @tags  微信
// @Description 微信支付
// @Summary 微信支付
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Token"
// @Param total_fee query int true "费用（分"
// @Param fee_type query string true "币种"
// @Param level query int true "公司等级"
// @Param month query int true "购买月数"
// @Param request_no query string true "公司编号"
// @Param open_id query string  true "微信openid"
// @Param pay_statement  query string true "商品名称"
// @Success 200 {string} json ""
// @Router /auth/pay/wxapp [post]
func WxAppPay(c *gin.Context) {
	var (
		err     error
		payInfo models.WXAppPayInfo
	)

	err = c.ShouldBindWith(&payInfo, binding.JSON)
	if err != nil {
		common.GenResponse(c, common.PAY_PARAMS_ERR, err, "failed")
		return
	}


	// 商户支付信息
	weChatInfo := common.WeChatInfo
	var reqMap = make(map[string]interface{}, 0)
	u2, err := uuid.NewV4()

	if err != nil {
		common.GenResponse(c, common.PAY_RAND_PARAM_ERR, err, "failed")
		return
	}
	var requestIp string
	remoteAdds := strings.Split(c.Request.RemoteAddr, ":")
	if len(remoteAdds) == 2 {
		requestIp = remoteAdds[0]
	}
	nonceStr := strings.Replace(u2.String(), "-", "", -1)
	orderNumber := fmt.Sprintf("%s-%d", time.Now().Format(common.TIME_FORMAT_ORDER), c.GetInt(common.LOGIN_USER_ID))
	reqMap["appid"] = weChatInfo.AppID           // 微信小程序appid
	reqMap["mch_id"] = weChatInfo.MchID          // 商户号
	reqMap["device_info"] = "WEB"                // 微信小程序
	reqMap["nonce_str"] = nonceStr               // 随机数
	reqMap["body"] = payInfo.PayStatement        // 商品描述
	reqMap["fee_type"] = payInfo.FeeType         // 币种
	reqMap["notify_url"] = weChatInfo.NotifyUrl  // 通知地址
	reqMap["openid"] = payInfo.OpenId            // 用户唯一标识
	reqMap["out_trade_no"] = orderNumber         // 订单号
	reqMap["spbill_create_ip"] = requestIp       // 用户端ip
	reqMap["total_fee"] = payInfo.TotalFee * 100 // 订单总金额，单位为分
	reqMap["trade_type"] = "JSAPI"               // trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识
	reqMap["sign"] = WxPayCalcSign(reqMap, weChatInfo.PaySignKey)
	reqStr := Map2Xml(reqMap)

	client := &http.Client{}
	req, err := http.NewRequest("POST", weChatInfo.PayUrl, strings.NewReader(reqStr))
	if err != nil {
		common.GenResponse(c, common.PAY_REQUEST_POST_ERR, err, "failed")
		return
	}
	req.Header.Set("Content-Type", "text/xml;charset=utf-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		common.GenResponse(c, common.PAY_ERR, err, "failed")
		return
	}
	var resp1 models.WXPayResp
	err = xml.Unmarshal(body2, &resp1)
	if err != nil {
		common.GenResponse(c, common.PAY_RESPONSE_UNMARSHAL_ERR, err, "failed")
		return
	}
	if resp1.PrepayId != "" {
		// 再次签名
		var resMap = make(map[string]interface{}, 0)
		resMap["appId"] = weChatInfo.AppID
		resMap["nonceStr"] = resp1.NonceStr                                  //商品描述
		resMap["package"] = common.StringsJoin("prepay_id=", resp1.PrepayId) //商户号
		resMap["signType"] = "MD5"                                           //签名类型
		resMap["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)       //当前时间戳

		resMap["paySign"] = WxPayCalcSign(resMap, weChatInfo.PaySignKey)
		// 写入数据库
	} else {
		common.GenResponse(c, common.PAY_RESPONSE_ERR, resp1, "failed")

	}
}

//WxPayCalcSign 微信支付计算签名的函数
func WxPayCalcSign(mReq map[string]interface{}, key string) (sign string) {
	//STEP 1, 对key进行升序排序.
	sortedKeys := make([]string, 0)
	for k, _ := range mReq {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	//STEP2, 对key=value的键值对用&连接起来，略过空值
	var signStrings string
	for _, k := range sortedKeys {

		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			signStrings = common.StringsJoin(signStrings, k, "=", value, "&")
		}
	}

	//STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = common.StringsJoin(signStrings, "key=", key)
	}
	//STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings)) //
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))

	return upperSign
}

//Map2Xml 微信支付计算签名的函数
func Map2Xml(mReq map[string]interface{}) (xml string) {
	sb := bytes.Buffer{}
	sb.WriteString("<xml>")
	for k, v := range mReq {
		vToStr := fmt.Sprintf("%v", v)
		sb.WriteString(common.StringsJoin("<", k, ">", vToStr, "</", k, ">"))
	}
	sb.WriteString("</xml>")
	return sb.String()
}

// @tags  微信
// @Description 微信支付回调
// @Summary 微信支付回调
// @Accept  json
// @Produce  json
// @Param record_id query int true "支付记录"
// @Param status query bool true "状态"
func WxAppPayAsyncBack(c *gin.Context) {
	type payResult struct {
		RecordId int  `json:"record_id" binding:"required"`
		Status   bool `json:"status" binding:"required"`
	}
	var (
		err    error
		result payResult
		code   int
	)

	err = c.ShouldBindWith(&result, binding.JSON)
	if err != nil {
		code = common.BINDING_JSON_ERR
		common.GenResponse(c, code, err, err.Error())
		return
	}
	//err, code = services.PayResult(result.RecordId, result.Status, c.GetInt(common.LOGIN_USER_ID))
	if err == nil {
		common.GenResponse(c, common.SUCCESSED, result, "success")
	} else {
		common.GenResponse(c, code, nil, err.Error())
	}
}
