package wxpay

import (
	"encoding/xml"
)

// PlaceOrderResult represent place order reponse message from weixin pay.
// For field explanation refer to: http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_1
type PlaceOrderResult struct {
	XMLName     xml.Name `xml:"xml"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	AppId       string   `xml:"appid"`
	MchId       string   `xml:"mch_id"`
	DeviceInfo  string   `xml:"device_info"`
	NonceStr    string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDesc string   `xml:"err_code_des"`
	TradeType   string   `xml:"trade_type"`
	PrepayId    string   `xml:"prepay_id"`
	CodeUrl     string   `xml:"code_url"`
}

func (this *PlaceOrderResult) ToMap() map[string]string {
	retMap, err := ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}

// Parse the reponse message from weixin pay to struct of PlaceOrderResult
func ParsePlaceOrderResult(resp []byte) (PlaceOrderResult, error) {
	placeOrderResult := PlaceOrderResult{}
	err := xml.Unmarshal(resp, &placeOrderResult)
	if err != nil {
		return placeOrderResult, err
	}

	return placeOrderResult, nil
}

// QueryOrder Result represent query response message from weixin pay
// Refer to https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7&index=8
type QueryOrderResult struct {
	XMLName            xml.Name `xml:"xml"`
	Appid              string   `xml:"appid"`
	Attach             string   `xml:"attach"`
	BankType           string   `xml:"bank_type"`
	CashFee            string   `xml:"cash_fee"`
	CashFeeType        string   `xml:"cash_fee_type"`
	CouponCount        string   `xml:"coupon_count"`
	CouponFee          string   `xml:"coupon_fee"`
	CouponFee0         string   `xml:"coupon_fee_0"`
	CouponFee1         string   `xml:"coupon_fee_1"`
	CouponFee2         string   `xml:"coupon_fee_2"`
	CouponFee3         string   `xml:"coupon_fee_3"`
	CouponFee4         string   `xml:"coupon_fee_4"`
	CouponFee5         string   `xml:"coupon_fee_5"`
	CouponFee6         string   `xml:"coupon_fee_6"`
	CouponFee7         string   `xml:"coupon_fee_7"`
	CouponFee8         string   `xml:"coupon_fee_8"`
	CouponFee9         string   `xml:"coupon_fee_9"`
	CouponID0          string   `xml:"coupon_id_0"`
	CouponID1          string   `xml:"coupon_id_1"`
	CouponID2          string   `xml:"coupon_id_2"`
	CouponID3          string   `xml:"coupon_id_3"`
	CouponID4          string   `xml:"coupon_id_4"`
	CouponID5          string   `xml:"coupon_id_5"`
	CouponID6          string   `xml:"coupon_id_6"`
	CouponID7          string   `xml:"coupon_id_7"`
	CouponID8          string   `xml:"coupon_id_8"`
	CouponID9          string   `xml:"coupon_id_9"`
	CouponType0        string   `xml:"coupon_type_0"`
	CouponType1        string   `xml:"coupon_type_1"`
	CouponType2        string   `xml:"coupon_type_2"`
	CouponType3        string   `xml:"coupon_type_3"`
	CouponType4        string   `xml:"coupon_type_4"`
	CouponType5        string   `xml:"coupon_type_5"`
	CouponType6        string   `xml:"coupon_type_6"`
	CouponType7        string   `xml:"coupon_type_7"`
	CouponType8        string   `xml:"coupon_type_8"`
	CouponType9        string   `xml:"coupon_type_9"`
	FeeType            string   `xml:"fee_type"`
	IsSubscribe        string   `xml:"is_subscribe"`
	MchID              string   `xml:"mch_id"`
	NonceStr           string   `xml:"nonce_str"`
	Openid             string   `xml:"openid"`
	OutTradeNo         string   `xml:"out_trade_no"`
	ResultCode         string   `xml:"result_code"`
	ReturnCode         string   `xml:"return_code"`
	SettlementTotalFee string   `xml:"settlement_total_fee"`
	Sign               string   `xml:"sign"`
	SignType           string   `xml:"sign_type"`
	SubMchID           string   `xml:"sub_mch_id"`
	TimeEnd            string   `xml:"time_end"`
	TotalFee           string   `xml:"total_fee"`
	TradeType          string   `xml:"trade_type"`
	TransactionID      string   `xml:"transaction_id"`
}

func (this *QueryOrderResult) ToMap() map[string]string {
	retMap, err := ToMap(this)
	if err != nil {
		panic(err)
	}

	return retMap
}

func ParseQueryOrderResult(resp []byte) (QueryOrderResult, error) {
	queryOrderResult := QueryOrderResult{}
	err := xml.Unmarshal(resp, &queryOrderResult)
	if err != nil {
		return queryOrderResult, err
	}

	return queryOrderResult, nil
}
