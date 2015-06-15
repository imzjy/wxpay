package wxpay

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// AppTrans is abstact of Transaction handler. With AppTrans, we can get prepay id
type AppTrans struct {
	Config *WxConfig
}

// Initialized the AppTrans with specific config
func NewAppTrans(cfg *WxConfig) *AppTrans {
	return &AppTrans{Config: cfg}
}

// Submit the order to weixin pay and return the prepay id if success,
// Prepay id is used for app to start a payment
// If fail, error is not nil, check error for more information
func (this *AppTrans) Submit(orderId string, amount float64, desc string, clientIp string) (string, error) {

	if this.Config.AppId == "" ||
		this.Config.MchId == "" ||
		this.Config.AppKey == "" ||
		this.Config.NotifyUrl == "" ||
		this.Config.NotifyUrl == "" ||
		this.Config.PlaceOrderUrl == "" ||
		this.Config.TradeType == "" {
		return "", errors.New("Please set key and cert")
	}

	odrInXml := this.orderInXmlString(orderId, fmt.Sprintf("%.0f", amount), desc, clientIp)
	resp, err := this.doRequest([]byte(odrInXml))
	if err != nil {
		return "", err
	}

	resultMsg, err := ParseResultMsg(resp)
	if err != nil {
		return "", err
	}

	if resultMsg.ReturnCode != "SUCCESS" {
		return "", errors.New(resultMsg.ReturnMsg)
	}

	if resultMsg.ResultCode != "SUCCESS" {
		return "", fmt.Errorf("resutl code:%s, result description:%s", resultMsg.ErrCode, resultMsg.ErrCodeDesc)
	}

	return resultMsg.PrepayId, nil

}

// doRequest post the order in xml format with a sign
func (this *AppTrans) doRequest(body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", this.Config.PlaceOrderUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respData, nil
}

func (this *AppTrans) buildPreSignOrder(orderId, amount, desc, clientIp string) map[string]string {
	param := make(map[string]string)
	param["appid"] = this.Config.AppId
	param["attach"] = "透传字段" //optional
	param["body"] = desc
	param["mch_id"] = this.Config.MchId
	param["nonce_str"] = strconv.FormatInt(time.Now().UnixNano(), 36)
	param["notify_url"] = this.Config.NotifyUrl
	param["out_trade_no"] = orderId
	param["spbill_create_ip"] = clientIp
	param["total_fee"] = amount
	param["trade_type"] = "APP"

	return param
}

func (this *AppTrans) orderInXmlString(orderId, amount, desc, clientIp string) string {
	preSignOrder := this.buildPreSignOrder(orderId, amount, desc, clientIp)
	preSignStr := SortAndConcat(preSignOrder)
	sign := Sign(preSignStr, this.Config.AppKey)
	fmt.Println(sign)

	preSignOrder["sign"] = sign

	xml := "<xml>"
	for k, v := range preSignOrder {
		xml = xml + fmt.Sprintf("<%s>%s</%s>", k, v, k)
	}
	xml = xml + "</xml>"

	return xml
}
