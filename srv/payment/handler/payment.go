package handler

import (
	"context"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/alipay"
	"github.com/mamachengcheng/12306/srv/payment/conf"
	"github.com/mamachengcheng/12306/srv/payment/proto/payment"
	"strconv"
)

type Payment struct {
	AlipayConfig *conf.AlipayConfig
}

func TradeAppPay(appId, privateKey, subject, outTradeNo, totalAmount, NotifyUrl string, isProd bool) (payParam string, err error) {
	client := alipay.NewClient(appId, privateKey, isProd)

	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(alipay.PKCS1).
		SetNotifyUrl(NotifyUrl)

	body := make(gopay.BodyMap)
	body.Set("subject", subject)
	body.Set("out_trade_no", outTradeNo)
	body.Set("total_amount", totalAmount)

	payParam, err = client.TradeAppPay(body)
	return payParam, err
}

func (p *Payment) Pay(ctx context.Context, in *payment.PayRequest, out *payment.PayResponse) error {

	payParam, err := TradeAppPay(p.AlipayConfig.AppId, p.AlipayConfig.PrivateKey, in.Subject, in.OutTradeNo, strconv.Itoa(int(in.TotalAmount)), p.AlipayConfig.NotifyUrl, p.AlipayConfig.IsProd)
	if err == nil {
		out.IsSuccess = true
		out.PayParam = payParam
	}
	return err
}

func TradeRefund(appId, privateKey, refundReason, outTradeNo, refundAmount string, isProd bool) (isSuccess bool, err error) {
	client := alipay.NewClient(appId, privateKey, isProd)

	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetPrivateKeyType(alipay.PKCS1)

	body := make(gopay.BodyMap)
	body.Set("refund_reason", refundReason)
	body.Set("out_trade_no", outTradeNo)
	body.Set("refund_amount", refundAmount)

	aliRsp, err := client.TradeRefund(body)

	if err == nil && aliRsp.Response.Code == "10000" && aliRsp.Response.FundChange == "Y" {
		isSuccess = true
	}

	return isSuccess, err
}

func (p *Payment) Refund(ctx context.Context, in *payment.RefundRequest, out *payment.RefundResponse) error {
	isSuccess, err := TradeRefund(p.AlipayConfig.AppId, p.AlipayConfig.PrivateKey, in.RefundReason, in.OutTradeNo, strconv.Itoa(int(in.RefundAmount)), p.AlipayConfig.IsProd)
	out.IsSuccess = isSuccess
	return err
}
