package wxpay

import (
	"encoding/xml"
)

// ResultMsg represent this reponse message from weixin pay.
// For field explanation refer to: http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_1
type ResultMsg struct {
	XMLName     xml.Name `xml:"xml"`
	ReturnCode  string   `xml:"return_code"`
	ReturnMsg   string   `xml:"return_msg"`
	AppId       string   `xml:"appid"`
	MchId       string   `xml:"mch_id"`
	DeviceInfo  string   `xml:"device_info"`
	Nonce       string   `xml:"nonce_str"`
	Sign        string   `xml:"sign"`
	ResultCode  string   `xml:"result_code"`
	ErrCode     string   `xml:"err_code"`
	ErrCodeDesc string   `xml:"err_code_des"`
	TradeType   string   `xml:"trade_type"`
	PrepayId    string   `xml:"prepay_id"`
	CodeUrl     string   `xml:"code_url"`
}


// Parse the reponse message from weixin pay to struct of ResultMsg
func ParseResultMsg(msg []byte) (ResultMsg, error) {
	resultMsg := ResultMsg{}
	err := xml.Unmarshal(msg, &resultMsg)
	if err != nil {
		return resultMsg, err
	}

	return resultMsg, nil
}
