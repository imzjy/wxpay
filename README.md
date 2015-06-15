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
	TradeType:     "APP",
}
appTrans := wxpay.NewAppTrans(cfg)

//获取prepay id
prepayId, err := appTrans.Submit("WOBHXLNSDFFALB7NLKN4FLVMPY", 1, "订单描述", "114.25.139.11")
if err != nil {
	panic(err)
}
fmt.Println(prepayId)

//生成手机端发起支付所需信息，参考：http://pay.weixin.qq.com/wiki/doc/api/app.php?chapter=9_12&index=2
payRequest := appTrans.BuildPaymentRequest(prepayId)
fmt.Println(payRequest)

```

# document

Please refer to [gowalker](https://gowalker.org/github.com/imzjy/wxpay)