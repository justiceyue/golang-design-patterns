package proxy

import "testing"

func TestProxy(t *testing.T) {
	alipaymentproxy := PaymentProxy{
		realPay: ali{},
	}
	t.Log(alipaymentproxy.Pay("aaa"))
	wechatpaymentproxy := PaymentProxy{
		realPay: wechat{},
	}
	t.Log(wechatpaymentproxy.Pay("bbb"))

	//use Factory
	apf := aliPaymentFac{}
	t.Log(Pay(apf, "ccc"))
	wpf := wechatPaymentFac{}
	t.Log(Pay(wpf, "ddd"))
}

func Pay(paymentFac PaymentFactory, order string) string {
	paymentProxy := PaymentProxy{
		realPay: paymentFac.Create(),
	}
	return paymentProxy.Pay(order)
}
