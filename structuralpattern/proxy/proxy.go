package proxy

import (
	"fmt"
)

type PaymentService interface {
	Pay(order string) string
}

type wechat struct {
}

func (wechat) Pay(order string) string {
	return fmt.Sprintf("wechat pay order:%s", order)
}

type ali struct {
}

func (ali) Pay(order string) string {
	return fmt.Sprintf("ali pay order:%s", order)
}

type PaymentFactory interface {
	Create() PaymentService
}

type aliPaymentFac struct {
}

func (aliPaymentFac) Create() PaymentService {
	return ali{}
}

type wechatPaymentFac struct {
}

func (wechatPaymentFac) Create() PaymentService {
	return wechat{}
}

type PaymentProxy struct {
	realPay PaymentService
}

func (p PaymentProxy) Pay(order string) string {
	fmt.Println("处理" + order)
	fmt.Println("1校验签名")
	fmt.Println("2格式化订单数据")
	fmt.Println("3参数检查")
	fmt.Println("4记录请求日志")
	return p.realPay.Pay(order)
}
