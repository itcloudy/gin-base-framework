package models

type WXAppPayInfo struct {
	FeeType      string `json:"fee_type"`      // 支付币种
	Level        int    `json:"level"`         // 公司级别
	TotalFee     int    `json:"total_fee"`     // 总费用
	OpenId       string `json:"open_id"`       // 用户openID
	PayStatement string `json:"pay_statement"` // 支付商品
	Month        int    `json:"month"`         // 购买月数
	RequestNo    string `json:"request_no"`    // 公司编号
}
