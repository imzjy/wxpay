# wxpay

Go语言微信App支付后台实现

Backend implementation of weixin pay(app) in golang 


# usage

```go
//初始化
cfg := &wxpay.WxConfig{
	AppId:         "应用程序Id, 从https://open.weixin.qq.com上可以看得到",
	AppKey:        "API密钥, 在 商户平台->账户设置->API安全 中设置",
	MchId:         "商户号",
	NotifyUrl:     "后台通知地址",
	PlaceOrderUrl: "https://api.mch.weixin.qq.com/pay/unifiedorder",
	QueryOrderUrl: "https://api.mch.weixin.qq.com/pay/orderquery",
	TradeType:     "APP",
}
appTrans, err := wxpay.NewAppTrans(cfg)
if err != nil {
	panic(err)
}

//获取prepay id，手机端得到prepay id后加上验证就可以使用这个id发起支付调用
prepayId, err := appTrans.Submit("WOBHXLNSDFFALB7NLKN4FLVMPY", 1, "订单描述", "114.25.139.11")
if err != nil {
	panic(err)
}
fmt.Println(prepayId)

//加上Sign，已方便手机直接调用
payRequest := appTrans.NewPaymentRequest(prepayId)
fmt.Println(payRequest)

//查询订单接口
queryResult, err := appTrans.Query("1008450740201411110005820873")
if err != nil {
	panic(err)
}
fmt.Println(queryResult)

```

# document

Please refer to [gowalker](https://gowalker.org/github.com/imzjy/wxpay)