package templatemethod

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

/*
模版方法是一种行为设计模式， 它在基类中定义了一个算法的框架， 允许子类在不修改结构的情况下重写算法的特定步骤。
模拟短信验证码发送过程
1.生成随机的验证code。
2.在缓存中保存这组数字以便进行后续验证。
3.准备内容。
4.发送通知。
同时假设保存验证码和发送验证消息的算法一致，会变的只有生成验证码和生成发送消息
*/

type VerificationMessage interface {
	GenRandomCode() string
	SaveCode(code string)
	GenMessage(code string) string
	SendMessage(message string)
}

type OTP struct {
	vfc VerificationMessage
}

//模板方法
func (o *OTP) SendVerificationMessageToUser() {
	code := o.vfc.GenRandomCode()
	o.vfc.SaveCode(code)
	message := o.vfc.GenMessage(code)
	o.vfc.SendMessage(message)
}

type OTPBase struct {
}

func (o *OTPBase) SaveCode(code string) {
	fmt.Println("保存随机码至缓存")
}

func (o *OTPBase) SendMessage(message string) {
	fmt.Printf("发送信息(%s)至用户\n", message)
}

type Email struct {
	OTPBase
}

func (e *Email) GenRandomCode() string {
	return uuid.NewString()
}

func (e *Email) GenMessage(code string) string {
	return fmt.Sprintf("email_%s", code)
}

type SMS struct {
	OTPBase
}

func (s *SMS) GenRandomCode() string {
	return fmt.Sprint(rand.Int63n(10000000))
}

func (s *SMS) GenMessage(code string) string {
	return fmt.Sprintf("sms_%s", code)
}
