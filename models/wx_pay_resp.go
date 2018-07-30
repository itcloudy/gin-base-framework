package models

//WXPayResp 响应信息
type WXPayResp struct {
	ReturnCode string `json:"return_code" xml:"return_code"`
	ReturnMsg  string `json:"return_msg" xml:"return_msg"`
	NonceStr   string `json:"nonce_str" xml:"nonce_str"`
	PrepayId   string `json:"prepay_id" xml:"prepay_id"`
}
